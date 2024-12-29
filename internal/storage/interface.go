package storage

type Storage interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string) bool
}

