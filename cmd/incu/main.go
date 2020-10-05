package main

import (
	"flag"
	"incu/pkg/dalet"
)

var (
	redisConn = flag.String("redis", "redis:6379", "redis connection string")
)

func main() {
	flag.Parse()
	bucket := dalet.NewMockBucket()

	source := dalet.NewNetworkSource("https://google.com")
	// source := dalet.NewStaticSource("./dalet.xml")
	w := dalet.NewWatcher(source, bucket)
	w.Start()

	// // rdb := redis.NewClient(&redis.Options{
	// // 	Addr:     *redisConn,
	// // 	Password: "",
	// // 	DB:       0,
	// // })

	// // bucket := dalet.NewRedisBucket(rdb)

}
