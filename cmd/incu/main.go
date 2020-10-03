package main

import (
	"flag"
	"incu/pkg/dalet"
	"os"
)

var (
	redisConn = flag.String("redis", "redis:6379", "redis connection string")
)

func main() {
	flag.Parse()
	f, _ := os.Open("./dalet.xml")
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     *redisConn,
	// 	Password: "",
	// 	DB:       0,
	// })

	// bucket := dalet.NewRedisBucket(rdb)
	bucket := dalet.NewMockBucket()
	w := dalet.NewWatcher(f, bucket)
	w.Start()

}
