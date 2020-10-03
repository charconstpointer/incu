package dalet

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"time"
)

type Watcher struct {
	source  io.Reader
	bucket  Bucket
	decoder *xml.Decoder
	Bm      BroadcastMonitor
}

func NewWatcher(r io.Reader, b Bucket) *Watcher {
	return &Watcher{
		source:  r,
		bucket:  b,
		decoder: xml.NewDecoder(r),
	}
}

func (w *Watcher) Start() {
	for {
		w.decoder.Decode(&w.Bm)

		t, err := time.Parse("2006-01-02T15:04:05", w.Bm.Current.StartTime)

		if err != nil {
			fmt.Println(err)
		}
		err = w.bucket.Save(context.Background(), w.Bm.Current)
		if err != nil {
			log.Fatal(err.Error())
		}
		dur := time.Duration(w.Bm.Current.Duration) * time.Millisecond
		end := t.Add(dur)
		fmt.Println("ends", end)
		time.Sleep(5 * time.Second)
	}
}
