## Beat Output rocketmq

This is a output implementation of [elastic beats](https://github.com/elastic/beats) for support [Filebeat](https://github.com/elastic/beats/tree/master/filebeat), [Metricbeat](https://github.com/elastic/beats/tree/master/metricbeat), [Functionbeat](https://github.com/elastic/beats/tree/master/x-pack/functionbeat), [Winlogbeat](https://github.com/elastic/beats/tree/master/winlogbeat), [Journalbeat](https://github.com/elastic/beats/tree/master/journalbeat), [Auditbeat](https://github.com/elastic/beats/tree/master/auditbeat) to [Apache rocketmq](https://github.com/apache/rocketmq)

### Compatibility
This output is developed and tested using Apache rocketmq-client-go  2.1.0 and Beats 7.9.1

### Download rocketmq-beat-output

```
mkdir -p $GOPATH/src/github.com/kukuzidian/
cd $GOPATH/src/github.com/kukuzidian/
git clone https://github.com/kukuzidian/rocketmq-beat-output
cd rocketmq-beat-output
```

### Build

#### Build Filebeat

Edit main.go file

```
package main

import (
    "os"
    _ "github.com/kukuzidian/rocketmq-beat-output/rocketmq"
    "github.com/elastic/beats/x-pack/filebeat/cmd"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

```
go build -o filebeat main.go
```

### Usage

#### Add following configuration to filebeat.yml
```
output.rocketmq:
  namesrvaddr: "172.16.14.128:9876"
  topic: beat_topic
  enabled: true
  codec.json:
    pretty: false
    escape_html: false
```
#### Start filebeat
```
./filebeat modules enable system
./filebeat modules list
./filebeat -c filebeat.yml -e
```

#### Build other beat

```
go build -o metricbeat metricbeat.go
go build -o filebeat filebeat.go
go build -o functionbeat functionbeat.go
go build -o journalbeat journalbeat.go
go build -o auditbeat auditbeat.go
go build -o winlogbeat winlogbeat.go
go build -o packetbeat packetbeat.go
```

### Configurations

#### Client
|Name|Description|Default|
|---|---|---|
|url| Configure the service URL for the rocketmq service |rocketmq://localhost:6650|
|certificate_path| path of tls cert file |""|
|private_key_path| path of tls key file |""|
|use_tls| Whether to turn on TLS, if to start, use protocol rocketmq+ssl |false|
|token| Access token information of cluster | "" |
|token_file_path| The file path where token is saved | "" |


#### Producer
|Name|Description|Default|
|---|---|---|
|topic| Specify the topic this producer will be publishing on. |""|
|name| Specify a name for the producer |""|
|send_timeout| Set the send timeout |30s|
|block_if_queue_full| Set whether the send and sendAsync operations should block when the outgoing message queue is full. |false|
|batching_max_messages| maximum number of messages in a batch |1000|
|batching_max_publish_delay| the batch delay |1ms|
|message_routing_mode| the message routing mode, SinglePartition,RoundRobinPartition, CustomPartition(0,1,2) |1|
|hashing_schema| JavaStringHash,Murmur3_32Hash(0,1) |0|
|compression_type| NONE,LZ4,ZLIB,ZSTD(0,1,2,3) |0|

### FAQ

