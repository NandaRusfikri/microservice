package kafka

import (
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	"service-user/constant"
	"service-user/dto"
	"time"
)

func NewKafkaProducer() *Producer {

	address, config := getKafkaConfig()
	CreateTopic(constant.TOPIC_NEW_ORDER)

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

var (
	group  = constant.SERVICE_NAME
	oldest = true
)

func getKafkaConfig() ([]string, *sarama.Config) {

	//version, err := sarama.ParseKafkaVersion(version)
	//if err != nil {
	//	log.Panicf("Error parsing Kafka version: %v", err)
	//}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Net.WriteTimeout = 5 * time.Second
	config.Producer.Retry.Max = 0
	//config.Version = version

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}

	if dto.CfgKafka.KafkaUser != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = dto.CfgKafka.KafkaUser
		config.Net.SASL.Password = dto.CfgKafka.KafkaPassword
	}
	return []string{dto.CfgKafka.KafkaAddress}, config
}
