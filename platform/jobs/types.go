package jobs

import "github.com/gocraft/work"

type WorkerContext struct{}

type JobEnqueuerInterface interface {
    Enqueue(jobName string, payload map[string]interface{}) (*work.Job, error)
}