package storage

import (
	grpc "dirScanner/internal/server/proto"
	"fmt"
)

type Cache struct {
	cacheMap map[string][]*grpc.File
}

func NewCacheStorage() *Cache {
	return &Cache{cacheMap: make(map[string][]*grpc.File)}
}

func (ch *Cache) SetValue(path string, files []*grpc.File) {
	ch.cacheMap[path] = files
}

func (ch *Cache) GetValue(path string) ([]*grpc.File, bool) {
	value, ok := ch.cacheMap[path]
	return value, ok
}

func (ch *Cache) Print(path string) {
	for _, val := range ch.cacheMap[path] {
		fmt.Println(val.FileName, val.Size, "Bytes")
	}
}
