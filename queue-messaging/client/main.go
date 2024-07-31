package main

import (
	"encoding/json"
	"queue-messaging/task"

	"github.com/hibiken/asynq"
)

func main() {
	redisOpt := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}

	client := asynq.NewClient(redisOpt)

	payload, _ := json.Marshal(task.SendEmail{
		From:    "5yTqG@example.com",
		To:      "5yTqG@example.com",
		Subject: "test",
		Content: "test",
	})
	task := asynq.NewTask(task.TaskSendEmail, payload)

	client.Enqueue(task)
}
