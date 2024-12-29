package test

import (
	"blockchain-parser/internal/storage"
	"testing"
)

func TestStorage_Set(t *testing.T) {
	storage := storage.NewMockStorage()

	storage.Set("key", "value")

	val, exists := storage.Get("key")
	if !exists {
		t.Fatalf("Failed to set value, key does not exist")
	}

	if val != "value" {
		t.Fatalf("Expected value 'value', got '%v'", val)
	}
}

func TestStorage_GetMissingKey(t *testing.T) {
	storage := storage.NewMockStorage()

	_, exists := storage.Get("missing_key")
	if exists {
		t.Fatalf("Key should not exist")
	}
}

func TestStorage_GetExistingKey(t *testing.T) {
	storage := storage.NewMockStorage()

	storage.Set("key", "value")

	val, exists := storage.Get("key")
	if !exists {
		t.Fatalf("Key should exist")
	}

	if val != "value" {
		t.Fatalf("Expected value 'value', got '%v'", val)
	}
}

func TestStorage_DeleteMissingKey(t *testing.T) {
	storage := storage.NewMockStorage()

	success := storage.Delete("missing_key")
	if success {
		t.Fatalf("Key should not exist, but delete returned success")
	}
}

func TestStorage_DeleteExistingKey(t *testing.T) {
	storage := storage.NewMockStorage()

	storage.Set("key", "value")

	success := storage.Delete("key")
	if !success {
		t.Fatalf("Key should exist, but delete returned failure")
	}

	_, exists := storage.Get("key")
	if exists {
		t.Fatalf("Key should not exist after deletion")
	}
}
