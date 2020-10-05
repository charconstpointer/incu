package main

import (
	"context"
	"flag"
	"incu/pkg/dalet"
	"time"
)

var (
	redisConn = flag.String("redis", "redis:6379", "redis connection string")
)

func main() {
	flag.Parse()
	//inmem
	mb := dalet.NewMockBucket()
	//redis
	_ = dalet.NewRedisBucket(*redisConn)

	//network
	_ = dalet.NewNetworkSource("https://google.com")
	//static
	source := dalet.NewStaticSource("./dalet.xml")

	w := dalet.NewWatcher(source, mb)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	w.Start(ctx)

}
