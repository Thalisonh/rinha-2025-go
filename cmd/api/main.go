package main

import (
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/thalisonh/rinha-go/internal/adapter/handler"
	"github.com/thalisonh/rinha-go/internal/domain/service"
	"github.com/thalisonh/rinha-go/pkg/streams/kafka"
	redisService "github.com/thalisonh/rinha-go/pkg/streams/redis"
	"gopkg.in/Shopify/sarama.v1"
)

var (
	producerInstance sarama.SyncProducer
	producerOnce     sync.Once
	producerErr      error
)

func main() {
	r := gin.Default()

	var sender service.StreamSender
	switch os.Getenv("STREAM_PROVIDER") {
	case "kafka":
		brokers := strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ",")
		topic := getEnv("KAFKA_TOPIC", "payments")

		producerOnce.Do(func() {
			config := sarama.NewConfig()
			config.Producer.Return.Successes = true
			producerInstance, producerErr = sarama.NewSyncProducer(brokers, config)
		})

		if producerErr != nil {
			panic(producerErr)
		}

		sender = kafka.NewKafkaStreamSender(producerInstance, topic)
	default:
		redisClient = redis.NewClient(&redis.Options{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       0,
		})

		sender = redisService.NewRedisStreamSender(redisClient)
	}

	svc := service.NewExampleService(sender)
	h := &handler.Handler{Service: svc}

	handler.RegisterRoutes(r, h)

	r.Run(":9999")
}

var redisClient *redis.Client

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
