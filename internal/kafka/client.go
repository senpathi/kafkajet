package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

type Client interface {
	sarama.Client
}

func NewClient() Client {
	conf := sarama.NewConfig()
	conf.Version = sarama.V3_2_0_0
	conf.Producer.Return.Errors = true             // this must be true for sync producer
	conf.Producer.Return.Successes = true          // this must be true for sync producer
	conf.Producer.RequiredAcks = sarama.WaitForAll // wait for all makes sure the reliability of the produced message

	cli, err := sarama.NewClient([]string{"localhost:9092"}, conf)
	if err != nil {
		log.Fatalln(err)
	}

	return cli
}
