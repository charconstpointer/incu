package dalet

import "encoding/xml"

type BroadcastMonitor struct {
	XMLName xml.Name `xml:"BroadcastMonitor"`
	Current Track    `xml:"Current"`
	Next    Next     `xml:"Next"`
}

type Track struct {
	XMLName   xml.Name `xml:"Current"`
	StartTime string   `xml:"startTime"`
	Title     string   `xml:"titleName"`
	Author    string   `xml:"artistName"`
	Duration  uint     `xml:"itemDurationMS"`
}

type Next struct {
	XMLName   xml.Name `xml:"Next"`
	StartTime string   `xml:"startTime"`
	Title     string   `xml:"titleName"`
	Author    string   `xml:"artistName"`
	Duration  uint     `xml:"itemDurationMS"`
}
