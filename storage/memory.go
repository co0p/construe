package storage

import (
	"errors"
	"fmt"
)

type DocumentNotFoundError struct {
	key string
}

func (e DocumentNotFoundError) Error() string {
	return fmt.Sprintf("'%s' not found", e.key)
}

type MemoryStorage struct {
	store map[string]Document
}

func NewMemoryStorage() MemoryStorage {
	store := MemoryStorage{
		store: make(map[string]Document),
	}
	return store
}

func (m MemoryStorage) Read(key string) (Document, error) {
	doc, found := m.store[key]
	if !found {
		return Document{}, DocumentNotFoundError{key}
	}
	return doc, nil
}

func (m MemoryStorage) Save(key string, document Document) error {
	if key == "" {
		return errors.New("invalid key provided")
	}
	m.store[key] = document
	return nil
}
