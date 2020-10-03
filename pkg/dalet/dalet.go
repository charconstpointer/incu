package dalet

import "encoding/xml"

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
