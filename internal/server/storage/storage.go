package storage

import (
	grpc "dirScanner/internal/server/proto"
	"fmt"
	"github.com/inhies/go-bytesize"
	"sync"
)

type Cache struct {
	mx       sync.RWMutex
	cacheMap map[string][]*grpc.File
}

func NewCacheStorage() *Cache {
	return &Cache{cacheMap: make(map[string][]*grpc.File)}
}

func (ch *Cache) SetValue(path string, files []*grpc.File) {
	ch.mx.Lock()
	defer ch.mx.Unlock()
	ch.cacheMap[path] = files
}

func (ch *Cache) GetValue(path string) ([]*grpc.File, bool) {
	ch.mx.RLock()
	defer ch.mx.RUnlock()
	value, ok := ch.cacheMap[path]
	return value, ok
}

func (ch *Cache) Print(path string) {
	for _, val := range ch.cacheMap[path] {
		fmt.Println(val.FileName, bytesize.New((float64(val.Size))))
	}
}
