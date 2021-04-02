package services

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

func (t *Target) isRedis(wg *sync.WaitGroup) {
	defer wg.Done()
	if t.Redis {
		is_redis := t.ScanPort(6379, t.Timeout)
		if is_redis {
			client := t.is_redis_check()
			pong, err := client.Ping().Result()
			if err == nil {
				if pong == "PONG" {
					fmt.Println("[+] redis ", t.Ip, pong, err)
				}
				// Output: PONG <nil>
			}
		}
	}
}

func (t *Target) is_redis_check() *redis.Client {
	ip := t.Ip + ":6379"
	client := redis.NewClient(&redis.Options{
		Addr: ip,
	})
	return client
}
