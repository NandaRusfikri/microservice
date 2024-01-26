package kafka

import (
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	"service-order/dto"
	"time"
)

func NewKafka() *Producer {

	address, config := getKafkaConfig()
	CreateTopic("sarama")

	producers, err := sarama.NewSyncProducer(address, config)

	if err != nil {
		log.Errorf("Unable to create kafka producer got error %v", err)
		//return
	}

	kafka := &Producer{
		Producer: producers,
	}

	return kafka
}

func CreateTopic(topic string) error {
	address, config := getKafkaConfig()
	admin, err := sarama.NewClusterAdmin(address, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
		return err
	}
	defer func() { _ = admin.Close() }()
	//err = admin.CreateTopic("topic.test.1", &sarama.TopicDetail{
	//	NumPartitions:     1,
	//	ReplicationFactor: 1,
	//}, false)

	if err := admin.CreatePartitions(topic, 3, [][]int32{}, false); err != nil {
		log.Errorf("Error while creating topic: ", err.Error())
		return err
	}
	return err
}

func getKafkaConfig() ([]string, *sarama.Config) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if dto.CfgKafka.KafkaUser != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = dto.CfgKafka.KafkaUser
		kafkaConfig.Net.SASL.Password = dto.CfgKafka.KafkaPassword
	}
	return []string{dto.CfgKafka.KafkaAddress}, kafkaConfig
}
