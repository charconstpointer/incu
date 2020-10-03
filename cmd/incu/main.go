package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

type BroadcastMonitor struct {
	XMLName xml.Name `xml:"BroadcastMonitor"`
	Current Current  `xml:"Current"`
}

type Current struct {
	XMLName   xml.Name `xml:"Current"`
	StartTime string   `xml:"startTime"`
	Title     string   `xml:"titleName"`
	Author    string   `xml:"artistName"`
	Duration  uint     `xml:"itemDurationMS"`
}

func main() {
	f, _ := os.Open("./dalet.xml")
	d := xml.NewDecoder(f)
	var bm BroadcastMonitor
	d.Decode(&bm)

	t, err := time.Parse("2006-01-02T15:04:05", bm.Current.StartTime)

	if err != nil {
		fmt.Println(err)
	}
	dur := time.Duration(bm.Current.Duration) * time.Millisecond
	end := t.Add(dur)
	fmt.Println("ends", end)

}
