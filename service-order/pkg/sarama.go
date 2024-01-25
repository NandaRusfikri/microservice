package pkg

import (
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	"service-order/dto"
	"time"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func NewKafka() *KafkaProducer {

	kafkaConfig := getKafkaConfig(dto.CfgKafka.KafkaUser, dto.CfgKafka.KafkaPassword)
	topik := "sarama"

	brokerAddrs := []string{dto.CfgKafka.KafkaAddress}
	admin, err := sarama.NewClusterAdmin(brokerAddrs, kafkaConfig)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()
	//err = admin.CreateTopic("topic.test.1", &sarama.TopicDetail{
	//	NumPartitions:     1,
	//	ReplicationFactor: 1,
	//}, false)

	admin.CreatePartitions(topik, 3, [][]int32{}, false)
	if err != nil {
		log.Fatal("Error while creating topic: ", err.Error())
	}

	producers, err := sarama.NewSyncProducer(brokerAddrs, kafkaConfig)

	if err != nil {
		log.Errorf("Unable to create kafka producer got error %v", err)
		return nil
	}
	defer func() {
		if err := producers.Close(); err != nil {
			log.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	return &KafkaProducer{
		Producer: producers,
	}
}

func (p *KafkaProducer) KirimPesan(topic, msg string, partition int32) error {

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

func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
