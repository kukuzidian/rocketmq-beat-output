package rocketmq

import (
	"github.com/elastic/beats/v7/libbeat/outputs/codec"
)

type rocketmqConfig struct {
	ProduceName string `config:"producename"`
	NameSrvAddr string `config:"namesrvaddr" validate:"required"`
	Topic string `config:"topic" validate:"required"`
	Retry uint `config:"retry"`
	MaxMessageSize uint32 `config:"maxmessagesize"`
	CompressMessageSize uint32 `config:"compressmessagesize"`
	Codec codec.Config `config:"codec"`
	GroupName string `config:"groupName"`
	Tag string `config:"tag"`
}


func defaultConfig() rocketmqConfig{
	return rocketmqConfig{
		ProduceName:         "default",
		NameSrvAddr:         "127.0.0.1:9876",
		Topic:               "default",
		Retry:               2,
		MaxMessageSize:      10 * 1024 * 1024,
		CompressMessageSize: 4 * 1024,
		GroupName: "defaultGroup",
		Tag: "",
	}
}



