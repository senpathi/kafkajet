package kafka

import (
	"errors"
	"fmt"
	"time"

	"github.com/Shopify/sarama"

	"github.com/senpathi/kafkajet/internal/domain"
	domainErr "github.com/senpathi/kafkajet/internal/errors"
)

type client struct {
	conf           *sarama.Config
	clusterDetails domain.Cluster
}

type Client interface {
	Topics() ([]string, error)
	CreateTopics(details []domain.TopicDetails) ([]string, error)
}

func NewClient(cluster domain.Cluster) Client {
	conf := sarama.NewConfig()
	conf.Version = sarama.V3_2_0_0
	conf.Producer.Return.Errors = true             // this must be true for sync producer
	conf.Producer.Return.Successes = true          // this must be true for sync producer
	conf.Producer.RequiredAcks = sarama.WaitForAll // wait for all makes sure the reliability of the produced message

	return &client{
		conf:           conf,
		clusterDetails: cluster,
	}
}

func (c *client) Topics() ([]string, error) {
	cli, err := sarama.NewClient(c.clusterDetails.Brokers, c.conf)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	return cli.Topics()
}

func (c *client) CreateTopics(details []domain.TopicDetails) ([]string, error) {
	cli, err := sarama.NewClient(c.clusterDetails.Brokers, c.conf)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctrl, err := cli.Controller()
	if err != nil {
		return nil, err
	}

	topics := make([]string, 0, len(details))

	topicDetails := make(map[string]*sarama.TopicDetail)
	for _, detail := range details {
		topicDetails[detail.Name] = &sarama.TopicDetail{
			NumPartitions:     detail.NumPartitions,
			ReplicationFactor: detail.ReplicationFactor,
		}

		topics = append(topics, detail.Name)
	}

	res, err := ctrl.CreateTopics(&sarama.CreateTopicsRequest{
		TopicDetails: topicDetails,
		Timeout:      time.Second * 30,
	})
	if err != nil {
		return nil, err
	}

	for k, v := range res.TopicErrors {
		if errors.Is(v.Err, sarama.ErrNoError) {
			continue
		}

		msg := fmt.Sprintf("%s. topic: [%s]", v.Error(), k)
		return nil, domainErr.Error{
			Err:     errors.New(msg),
			Code:    domainErr.InvalidRequestErrorCode,
			Message: msg,
		}
	}

	return topics, err
}
