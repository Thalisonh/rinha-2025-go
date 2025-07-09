package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	redisService "github.com/thalisonh/rinha-go/pkg/streams/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       0,
	})

	clientRedis := redisService.NewRedisStreamSender(redisClient)
	for {
		resp, err := clientRedis.Get()
		if err != nil && err != redis.Nil {
			log.Printf("Erro ao ler do stream: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		result := []RedisStreamResult{}
		json.Unmarshal([]byte(resp), &result)

		for _, stream := range result {
			for _, msg := range stream.Messages {
				payloadStr, ok := msg.Values["payload"].(string)
				if !ok {
					continue
				}
				var payment CreatePaymentInput
				err := json.Unmarshal([]byte(payloadStr), &payment)
				if err != nil {
					continue
				}

				input := &CreatePaymentRequest{
					Amount:        payment.Amount,
					CorrelationID: payment.CorrelationID,
					RequestedAt:   time.Now().Format(time.RFC3339),
				}

				err = callEndpoint("http://localhost:8001/payments", input)
				if err != nil {
					callEndpoint("http://localhost:8002/payments", input)
					log.Printf("Erro ao chamar endpoint: %v", err)
					return
				}

				err = clientRedis.Delete(msg.ID)
				if err != nil {
					log.Printf("Erro ao deletar mensagem: %v", err)
					return
				}
				fmt.Printf("ID %s deletado", msg.ID)

				return
			}
		}
	}
}

type CreatePaymentRequest struct {
	Amount        float64 `json:"Amount"`
	CorrelationID string  `json:"CorrelationID"`
	RequestedAt   string  `json:"requestedAt"`
}

type CreatePaymentInput struct {
	Amount        float64 `json:"Amount"`
	CorrelationID string  `json:"CorrelationID"`
}

type RedisStreamMessage struct {
	ID     string                 `json:"ID"`
	Values map[string]interface{} `json:"Values"`
}

type RedisStreamResult struct {
	Stream   string               `json:"Stream"`
	Messages []RedisStreamMessage `json:"Messages"`
}

func callEndpoint(url string, payload interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("erro ao serializar payload: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("erro ao chamar %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("resposta de %s: %s (status %d)", url, string(body), resp.StatusCode)
	}

	log.Printf("Resposta de %s: %s", url, string(body))
	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
