package sink

import (
	"fmt"
	//"github.com/go-redis/redis"
)

type Redis struct {
	URL string
}

func NewRedis(url string) Redis {
	return Redis{URL: url}
}

func (r Redis) Push(payload []byte) error {
	fmt.Println("redis: ", string(payload), " at ", r.URL)
	//TODO PUSH
	return nil
}
