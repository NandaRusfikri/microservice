package kafka

import (
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

type Producer struct {
	Producer sarama.SyncProducer
}

func (p *Producer) SendMessage(topic, msg string, partition int32) error {

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
		//Partition: partition,
	}

	partition, offset, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		log.Errorf("Send message error: %v", err)
		return err
	}

	log.Infof("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return nil
}
