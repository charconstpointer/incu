package dalet

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

type Watcher struct {
	source Source
	bucket Bucket
	Bm     BroadcastMonitor
}

func NewWatcher(s Source, b Bucket) *Watcher {
	return &Watcher{
		source: s,
		bucket: b,
	}
}

func (w *Watcher) Start(c context.Context) {
	go w.process(c)
	select {
	case <-c.Done():
		break
	}
}

func (w *Watcher) process(c context.Context) {
	for {
		select {
		case <-c.Done():
			fmt.Printf("Done\n")
			break
		default:
			t, err := w.source.GetTrack()
			if err != nil {
				log.Printf("Could not fetch next track")
				time.Sleep(1 * time.Minute)
				continue
			}
			d := xml.NewDecoder(t)
			err = d.Decode(&w.Bm)
			if err != nil {
				log.Print("Could not decode the reader\n")
				time.Sleep(1 * time.Minute)
				continue
			}
			date, err := time.Parse("2006-01-02T15:04:05", w.Bm.Current.StartTime)

			if err != nil {
				fmt.Println(err)
			}
			err = w.bucket.Save(context.Background(), w.Bm.Current)
			if err != nil {
				log.Fatal(err.Error())
			}
			dur := time.Duration(w.Bm.Current.Duration) * time.Millisecond
			end := date.Add(dur)
			fmt.Println("ends", end)
			sleepFor := end.Sub(time.Now())
			log.Println("sleep for ", sleepFor)
			time.Sleep(sleepFor)
		}
	}
}
