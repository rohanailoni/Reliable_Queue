package Queue

import (
	"ReliableQueue/Output"
	Interface "ReliableQueue/Queue/Interface"
	"ReliableQueue/model"
	"context"
)

type QueueClient struct {
	Queue Interface.IqueueConnectionInterface
}

func (que *QueueClient) Connect(_addr string, _password string, _db int) *model.MessageBroker {
	return que.Queue.Connect(_addr, _password, _db)
}

func (que *QueueClient) LPUSH(ctx context.Context, queueName string, event *model.Event) error {
	return que.Queue.LPUSH(ctx, queueName, event)
}
func (que *QueueClient) RPOP(ctx context.Context, queueName string) *Output.EventStatus {
	return que.Queue.RPOP(ctx, queueName)
}
func (que *QueueClient) ZADD(ctx context.Context, setName string, score int64, members []byte) error {
	return que.Queue.ZADD(ctx, setName, score, members)
}
