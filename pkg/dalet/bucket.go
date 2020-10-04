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

func NewRedisBucket(c *redis.Client) *RedisBucket {
	return &RedisBucket{
		conn: c,
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
