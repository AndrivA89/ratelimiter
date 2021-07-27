package ratelimiter

import (
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

const rateLimit = time.Minute * 1

var (
	maxTasksInMinute    = int32(10) //default
	maxTasksInMinuteInt int
	maxWorkers          = 5 //default
	exit                = make(chan struct{})
	rate                = time.Tick(rateLimit)
	tasks               int32
	err                 error
)

func RateLimitCall(c chan int) {
	if value, ok := os.LookupEnv("RATE_LIMITER_MAX_TASKS_IN_MINUTE"); ok {
		maxTasksInMinuteInt, err = strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
			return
		}
		maxTasksInMinute = int32(maxTasksInMinuteInt)
	}

	if value, ok := os.LookupEnv("RATE_LIMITER_MAX_WORKERS"); ok {
		maxWorkers, err = strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	lenChan := len(c)

	go func() {
		for {
			select {
			case <-rate:
				tasks = 0
			case <-exit:
				os.Exit(0)
			}
		}
	}()

	for len(c) > 0 {
		for i := 0; i < maxWorkers; i++ {
			go func(i int) {
				if tasks < maxTasksInMinute {
					lenChan--
					if lenChan < 0 {
						exit <- struct{}{}
					}
					atomic.AddInt32(&tasks, 1)
					number := <-c
					fmt.Printf("worker %v, task %v\n", i+1, number)
				}
			}(i)
		}
	}
}
