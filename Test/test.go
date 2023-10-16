package Test

import (
	que "ReliableQueue/Queue"
	"ReliableQueue/Queue/structs"
	"ReliableQueue/model"
	"context"
	"github.com/google/uuid"
	"time"
)

const (
	testQueue = "test-queue"
)

func main() {
	redisClient := &structs.RedisClient{}
	queue := "test-queue"
	ctx := context.Background()
	client := &que.QueueClient{Queue: redisClient}
	client.Connect("localhost:6379", "", 0)
	event := &model.Event{
		EventId:          uuid.New().String(),
		SubmissionId:     uuid.New().String(),
		Runtime:          5,
		CodeBucket:       "S3:<URL>",
		FirstEventEpoch:  time.Now().UnixMilli(),
		LatestEventEpoch: time.Now().UnixMilli(),
		Retries:          0,
	}

}
