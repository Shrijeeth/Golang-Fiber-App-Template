package jobs

import (
	"fmt"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"os"
	"strconv"
	"time"
)

type WorkerContext struct{}

var (
	WorkerPool   *work.WorkerPool
	RedisJobPool *redis.Pool
	JobEnqueuer  *work.Enqueuer
)

func InitWorkerPool() error {
	RedisJobPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   50,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
			if err != nil {
				conn.Close()
				return nil, err
			}

			if os.Getenv("REDIS_PASSWORD") != "" {
				if _, err := conn.Do("AUTH", os.Getenv("REDIS_PASSWORD")); err != nil {
					conn.Close()
					return nil, err
				}
			}

			if _, err := conn.Do("SELECT", os.Getenv("REDIS_DB_NUMBER")); err != nil {
				conn.Close()
				return nil, err
			}

			return conn, nil
		},
	}

	WorkerPool = work.NewWorkerPool(WorkerContext{}, 10, os.Getenv("APP_NAMESPACE"), RedisJobPool)
	WorkerPool.Middleware(JobLogger)

	err := RegisterJobs()
	if err != nil {
		return err
	}

	JobEnqueuer = work.NewEnqueuer(os.Getenv("APP_NAMESPACE"), RedisJobPool)

	WorkerPool.Start()

	return nil
}

func CloseWorkerPool() error {
	WorkerPool.Stop()
	err := RedisJobPool.Close()
	return err
}

func RegisterJobs() error {
	WorkerPool.JobWithOptions(SampleMailJobName, work.JobOptions{
		Priority:       10,
		MaxFails:       1,
		MaxConcurrency: 3,
	}, SampleMailJob)

	return nil
}

func JobLogger(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	fmt.Println("Job Params: ", job.Args)
	return next()
}

func IsJobWorkerRequired() bool {
	isJobWorkerRequired, _ := strconv.Atoi(os.Getenv("JOB_WORKER_REQUIRED"))
	return isJobWorkerRequired == 1
}
