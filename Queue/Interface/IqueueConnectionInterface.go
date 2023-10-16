package Interface

import (
	"ReliableQueue/Output"
	"ReliableQueue/model"
	"context"
)

type IqueueConnectionInterface interface {
	Connect(_addr string, _password string, _db int) *model.MessageBroker

	LPUSH(ctx context.Context, queueName string, event *model.Event) error

	RPOP(ctx context.Context, queueName string) *Output.EventStatus

	ZADD(ctx context.Context, setName string, score int64, members []byte) error
}
