package redis

import (
	"context"
	"encoding/json"

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

func (r *RedisStreamSender) Get() (string, error) {
	ctx := context.Background()
	res, err := r.client.XRead(ctx, &redis.XReadArgs{
		Streams: []string{"payments", "0"},
		Count:   100,
		Block:   300,
	}).Result()
	if err != nil {
		return "", err
	}

	json, _ := json.Marshal(res)
	return string(json), nil
}

func (r *RedisStreamSender) Delete(id string) error {
	ctx := context.Background()

	_, err := r.client.XDel(ctx, "payments", id).Result()
	if err != nil {
		return err
	}

	return nil
}
