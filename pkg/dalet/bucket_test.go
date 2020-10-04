package dalet_test

import (
	"context"
	"incu/pkg/dalet"
	"testing"
)

func TestBucketSave(t *testing.T) {
	b := dalet.NewMockBucket()
	c := dalet.Track{
		Author: "author",
	}
	b.Save(context.Background(), c)
	got := b.Head()
	if got != c {
		t.Errorf("Expeced %v, got %v", c, b)
	}
	c = dalet.Track{
		Author: "author2",
	}
	b.Save(context.Background(), c)
	got = b.Head()
	if got != c {
		t.Errorf("Expeced %v, got %v", c, b)
	}
	if got.Author != "author2" {
		t.Errorf("Expected author to be %s, but instead got %s", "author2", got.Author)
	}
}
