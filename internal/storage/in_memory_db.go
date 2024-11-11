package storage

import "sync"

type InMemoryDB struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{data: make(map[string]string)}
}

func (db* InMemoryDB) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

func (db* InMemoryDB) Get(key string) (string, bool){
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, exists := db.data[key]
	return value, exists
}