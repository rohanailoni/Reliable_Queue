package structs

import (
	Output "ReliableQueue/Output"
	"ReliableQueue/Queue/serialization"
	"ReliableQueue/model"
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisClient struct {
	client *redis.Client
}

func (r *RedisClient) Connect(_addr string, _password string, _db int) *model.MessageBroker {
	rdb := redis.NewClient(&redis.Options{
		Addr:     _addr,
		Password: _password,
		DB:       _db,
	})
	r.client = rdb
	return &model.MessageBroker{
		RedisClient: rdb,
	}
}

// List Operations
func (r *RedisClient) LPUSH(ctx context.Context, queueName string, event *model.Event) error {
	data, err := serialization.EventEncoder(event).Result()
	err = r.client.LPush(ctx, queueName, data).Err()
	return err
}

func (r *RedisClient) RPOP(ctx context.Context, queueName string) *Output.EventStatus {
	pop, err := r.client.RPop(ctx, queueName).Bytes()
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred: and this is recovered stage for error and QueueName", queueName, err)
		}
	}()
	event, err := serialization.EventDecoder(pop)
	return &Output.EventStatus{
		Base: Output.Base{
			Ctx:   ctx,
			Error: err,
		},
		Event: event,
	}

}

//Sorted Set Operations

/*
*
Adds the elements to the sorted set with one elements with score and current time and memeber as event bytes.
complexity O(log(N)).

returns an integer that specify number of memebers inserted into set but we have nothing to do with that so we do return only error.
*
*/
func (r *RedisClient) ZADD(ctx context.Context, setName string, score int64, members []byte) error {
	_, err := r.client.ZAdd(ctx, setName, &redis.Z{
		Score:  float64(score),
		Member: members,
	}).Result()
	//SortedSet Insertion Failure
	return err
}

//func (r *RedisClient) ZREM
