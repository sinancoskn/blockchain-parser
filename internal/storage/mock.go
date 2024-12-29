package storage

import (
	"sync"
)

type MockStorage struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewMockStorage() Storage {
	return &MockStorage{
		data: make(map[string]interface{}),
	}
}

func (m *MockStorage) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}

func (m *MockStorage) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, exists := m.data[key]
	return val, exists
}

func (m *MockStorage) Delete(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.data[key]; exists {
		delete(m.data, key)
		return true
	}

	return false
}
