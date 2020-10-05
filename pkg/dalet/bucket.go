package dalet

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Bucket interface {
	Save(ctx context.Context, c Track) error
}

type RedisBucket struct {
	conn *redis.Client
}

func NewRedisBucket(connString string) *RedisBucket {
	rdb := redis.NewClient(&redis.Options{
		Addr:     connString,
		Password: "",
		DB:       0,
	})
	return &RedisBucket{
		conn: rdb,
	}
}

func (b *RedisBucket) Save(ctx context.Context, c Track) error {
	err := b.conn.Set(ctx, "key", "value", 0).Err()
	return err
}

type MockBucket struct {
	head Track
}

func NewMockBucket() *MockBucket {
	return &MockBucket{}
}

func (b *MockBucket) Save(ctx context.Context, c Track) error {
	log.Printf("bucket => %s", c.Title)
	b.head = c
	return nil
}

func (b *MockBucket) Head() Track {
	return b.head
}
