package client

import (
	"sync"
)

var (
	access_token string
	muStrava     sync.RWMutex
)

func SetAccessToken(c string) {
	muStrava.Lock()
	defer muStrava.Unlock()
	access_token = c
}

func GetAccessToken() string {
	muStrava.RLock()
	defer muStrava.RUnlock()
	return access_token
}
