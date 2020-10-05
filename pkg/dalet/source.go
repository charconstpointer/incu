package dalet

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Source interface {
	GetTrack() (io.Reader, error)
}
type StaticSource struct {
	filePath string
}

type NetworkSource struct {
	client http.Client
	url    url.URL
}

func NewNetworkSource(address string) *NetworkSource {
	p, err := url.Parse(address)
	if err != nil {
		log.Fatalf("%s is not a valid URL\n", address)
	}

	return &NetworkSource{
		client: http.Client{},
		url:    *p,
	}
}

func (s *NetworkSource) GetTrack() (io.Reader, error) {
	r, err := http.NewRequest("GET", s.url.String(), nil)
	if r == nil {
		return nil, err
	}
	res, err := s.client.Do(r)
	if res == nil{
		return nil, err
	}
	return res.Body, err
}

func (s *StaticSource) GetTrack() (io.Reader, error) {
	f, err := os.Open(s.filePath)
	return f, err
}

func NewStaticSource(filePath string) *StaticSource {
	return &StaticSource{
		filePath: filePath,
	}
}
