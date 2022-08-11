package main

import (
	"time"

	"github.com/felixls/ratelimiter"
)

func main() {
	r, err := ratelimiter.NewThrottleRateLimiter(&ratelimiter.Config{
		Throttle: 1 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	ratelimiter.DoWork(r, 10)
}
