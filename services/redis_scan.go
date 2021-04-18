package services

import (
	"fmt"
	loggers "samalas/logger"
	"sync"

	"github.com/go-redis/redis"
)

func (t *Target) isRedis(wg *sync.WaitGroup) {
	defer wg.Done()
	if t.Redis {
		// - check redis port is open or not
		is_redis := t.ScanPort(6379, t.Timeout)
		msg := fmt.Sprintf("Scan redis port on IP : %s", t.Ip)
		loggers.SetLogger("info", msg)
		// - if redis is open
		if is_redis {
			client := t.is_redis_check()
			pong, err := client.Ping().Result()
			if err == nil {
				if pong == "PONG" {
					fmt.Println("[+] redis open", t.Ip, pong, "message", err)
				}
				// Output: PONG <nil>
			}
		}
	}
}

func (t *Target) is_redis_check() *redis.Client {
	ip := fmt.Sprintf("%s:6379", t.Ip)

	client := redis.NewClient(&redis.Options{
		Addr: ip,
	})
	return client
}
