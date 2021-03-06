package rocketmq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"strconv"
	"time"

	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/outputs/codec"
	"github.com/elastic/beats/v7/libbeat/publisher"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type rocketmqcli struct {
	rmqproducer rocketmq.Producer
	producename string
	namesrvaddr string
	topic string
	retry uint
	maxmessagesize uint32
	compressmessagesize uint32
	beat     beat.Info
	observer outputs.Observer
	codec    codec.Codec
	log      *logp.Logger
	tag string
}

func init()  {
	outputs.RegisterType("rocketmq",makeRocketmq)
}

func (r *rocketmqcli) Close() error {
	err := r.rmqproducer.Shutdown()
	return err
}

func (r *rocketmqcli) Publish(ctx context.Context, batch publisher.Batch) error {
	defer batch.ACK()
	events := batch.Events()
	r.observer.NewBatch(len(events))
	r.log.Warn("rocketmq", "Rocketmq received events: %d", len(events))
	cnt :=1000
	var slicearr = make([]string,1,1)
	var tmpslice = make([]*primitive.Message,0,cnt)
	dropped := 0
	for i := range events{
		event := &events[i]
		serializedEvent, err := r.codec.Encode(r.beat.Beat, &event.Content)

		if err != nil {
			if event.Guaranteed() {
				r.log.Error("Failed to serialize the event: %+v", err)
			} else {
				r.log.Error("Failed to serialize the event: %+v", err)
			}
			r.log.Error("Failed event: %v", event)

			dropped++
			continue
		}

		msg := &primitive.Message{
			Topic: r.topic,
			Body:  []byte(serializedEvent),
		}
		msg = msg.WithTag(r.tag)
		//msg = msg.WithShardingKey(time.Now().String())
		slicearr[0] = strconv.FormatInt(time.Now().UnixNano()/1000,10)
		msg = msg.WithKeys(slicearr)

		tmpslice = append(tmpslice, msg)
		if len(tmpslice) == cnt {
			res, err := r.rmqproducer.SendSync(context.Background(), tmpslice...)
			if err != nil {
				r.log.Error("send message error: %s\n", err)
			} else {
				r.log.Debug("send message success: result=%s\n", res.String())
			}
			//tmpslice = make([]*primitive.Message,0,cnt)
			tmpslice = tmpslice[0:0]
		}

	}
	if len(tmpslice) > 0 {
		res, err := r.rmqproducer.SendSync(context.Background(), tmpslice...)
		if err != nil {
			r.log.Error("send message error: %s\n", err)
		} else {
			r.log.Debug("send message success: result=%s\n", res.String())
		}
	}
	r.observer.Dropped(dropped)
	r.observer.Acked(len(events) - dropped)
	return nil
}

func (r *rocketmqcli) String() string {
	//panic("implement me")
	return "rocketmq topic="+r.topic
}

//func cleardata(t []*primitive.Message) []*primitive.Message {
//	for i:=0 ; i < len(t) ; i++ {
//		t[i] = nil
//	}
//	return t
//}

func makeRocketmq(
	_ outputs.IndexManager,
	beat beat.Info,
	observer outputs.Observer,
	cfg *common.Config,
) (outputs.Group, error) {
	config := defaultConfig()
	logp.Info("initialize rocketmq output")
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}
	logp.Info("init client %v", config)
	rmq := &rocketmqcli{
		namesrvaddr: config.NameSrvAddr,
		topic: config.Topic,
		beat:     beat,
		observer: observer,
	}
	if rmq == nil {
		logp.Err("========================rmq = null error ")
	}

	var err2 error
	err2 = rmq.initConfig(beat,config)
	if err2 != nil {
		rmq.log.Error("---initConfig error !")
	}
	if rmq.log == nil {
		logp.Err("========================rmq.log = null error ")
	}
	rmq.log.Info("---initConfig  success! ")

	err2 = rmq.rmqproducer.Start()
	if err2 != nil {
		rmq.log.Info("---rmqproducer started error! ")
	}else{
		rmq.log.Info("---rmqproducer started success! ")
	}
	return outputs.Success(config.BatchSize,config.Retry,rmq)
}

func (r *rocketmqcli) initConfig(info beat.Info,c rocketmqConfig) error {
	var err error
	r.log = logp.NewLogger("file")
	r.codec , err = codec.CreateEncoder(info,c.Codec)
	if err != nil {
		r.log.Error("--create CreateEncoder error!")
		return err
	}

	r.maxmessagesize = c.MaxMessageSize
	r.compressmessagesize = c.CompressMessageSize
	r.tag = c.Tag

	nsraddr , err := primitive.NewNamesrvAddr(c.NameSrvAddr)
	if err != nil {
		r.log.Error("rocketmq NameSrvAddr init error !")
		return err
	}

	r.rmqproducer, err = rocketmq.NewProducer(
		//producer.WithNsResovler(primitive.NewPassthroughResolver([]string{c.NameSrvAddr})),
		producer.WithNameServer(nsraddr),
		producer.WithRetry(2),
		producer.WithCreateTopicKey(c.Topic),
		producer.WithGroupName(c.GroupName),
	)

	if err != nil {
		r.log.Error("--create rmqproducer error!")
		return err
	}
	return nil
}