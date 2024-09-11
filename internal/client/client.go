package client

import (
	"sync"

	"github.com/zmb3/spotify/v2"
)

var (
	client *spotify.Client
	mu     sync.RWMutex
)

func SetClient(c *spotify.Client) {
	mu.Lock()
	defer mu.Unlock()
	client = c
}

func GetClient() *spotify.Client {
	mu.RLock()
	defer mu.RUnlock()
	return client
}
