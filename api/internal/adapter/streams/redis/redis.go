package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStreamSender struct {
	client *redis.Client
}

func NewRedisStreamSender(client *redis.Client) *RedisStreamSender {
	return &RedisStreamSender{client: client}
}

func (r *RedisStreamSender) Send(payload []byte) error {
	ctx := context.Background()
	_, err := r.client.XAdd(ctx, &redis.XAddArgs{
		Stream: "payments",
		Values: map[string]interface{}{
			"payload": string(payload),
		},
	}).Result()
	return err
}
