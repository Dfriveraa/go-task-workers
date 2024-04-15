package main

import (
	"github.com/dfriveraa/go-task-workers/asyncq/worker/tasks"
	"github.com/hibiken/asynq"
	"log"
)

// workers.go
func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379", Password: "password123"},
		asynq.Config{Concurrency: 10, Queues: map[string]int{"default": 9, "critical": 1}},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeWelcomeEmail, tasks.HandleWelcomeEmailTask)
	mux.HandleFunc(tasks.TypeReminderEmail, tasks.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
