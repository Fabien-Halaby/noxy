package model

import (
	"io"
	"net/http"
	"net/url"
	"sync"
)

type CacheEntry struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

type Cache struct {
	Store map[string]CacheEntry
	Mutex sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		Store: make(map[string]CacheEntry),
		Mutex: sync.RWMutex{},
	}
}

func (c *Cache) Get(key string) (CacheEntry, bool) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	entry, exists := c.Store[key]

	return entry, exists
}

func (c *Cache) Set(key string, entry CacheEntry) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Store[key] = entry
}

func (c *Cache) Clear() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Store = make(map[string]CacheEntry)
}

func ProxyRequest(client *http.Client, originURL *url.URL, r *http.Request) (CacheEntry, error) {
	targetURL := *originURL
	targetURL.Path = r.URL.Path
	targetURL.RawQuery = r.URL.RawQuery

	proxyReq, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
	if err != nil {
		return CacheEntry{}, nil
	}

	for key, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	resp, err := client.Do(proxyReq)
	if err != nil {
		return CacheEntry{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CacheEntry{}, err
	}

	return CacheEntry{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header.Clone(),
		Body:       body,
	}, nil
}
