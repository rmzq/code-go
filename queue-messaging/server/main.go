package main

import (
	"context"
	"encoding/json"
	"log"
	"queue-messaging/task"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	resiOpt := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}

	svr := asynq.NewServer(resiOpt, asynq.Config{
		Concurrency: 10,
	})

	mux := asynq.NewServeMux()

	mux.HandleFunc(task.TaskSendEmail, func(ctx context.Context, t *asynq.Task) error {
		// var payload task.SendEmail
		var payload2 map[string]any
		json.Unmarshal(t.Payload(), &payload2)
		log.Println("received payload: ", payload2["from"])
		log.Println("sending email...", payload2)
		time.Sleep(5 * time.Second)
		return nil
	})

	mux.HandleFunc(task.TaskUploadFile, func(ctx context.Context, t *asynq.Task) error {
		log.Println("uploading file...")
		return nil
	})

	if err := svr.Run(mux); err != nil {
		log.Fatal(err)
	}
}
