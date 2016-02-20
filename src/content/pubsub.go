package content

import (
	redis "gopkg.in/redis.v3"
)

func subscribe() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pb, err := client.Subscribe("gitpush")
	print(pb, err)
}
