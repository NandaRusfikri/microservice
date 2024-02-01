package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

type Producer struct {
	Producer sarama.SyncProducer
}

func (p *Producer) SendMessage(topic string, msg map[string]interface{}, partition int32) error {

	jsonByte, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(jsonByte),
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
