package main

import (
	que "ReliableQueue/Queue"
	"ReliableQueue/Queue/serialization"
	structs "ReliableQueue/Queue/structs"
	_const "ReliableQueue/const"
	"ReliableQueue/model"
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	redisClient := &structs.RedisClient{}
	ctx := context.Background()
	client := &que.QueueClient{Queue: redisClient}
	cc := client.Connect("localhost:6379", "", 0)

	defer cc.RedisClient.Close()

	//Pop an event from the Queue.
	event, err := client.RPOP(ctx, _const.MAIN_QUEUE).Result()
	if err != nil {

		//error login should be done<TBD>
	}

	// This is the Failure point where if the worker is lost or crashed here then there is no way of recovering the event.
	//This is quite sad part of redis where it is not considered as reliable.
	data, err := serialization.EventEncoder(&event).Result()
	if err != nil {
		//serialization error. Handling <TBD>
		fmt.Println("Handle serialization error")

	}

	err = client.ZADD(ctx, _const.TEST_QUEUE, time.Now().UnixMicro(), data)
	if err != nil {
		fmt.Println("Hanlde Insertion Error")
		//TBD
	}
	result := Process(event)
	if result {

	}

}

// This is temporarary function
func Process(event model.Event) bool {
	time.Sleep(time.Second * 10)
	n, _ := strconv.ParseInt(event.EventId, 10, 64)
	if n%2 == 0 {
		return true
	}
	return false

}
