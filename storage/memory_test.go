package storage

import "testing"

func TestNewMemoryStorage(t *testing.T) {
	store := NewMemoryStorage()

	_, err := store.Read("")
	if err == nil {
		t.Error("Expected error, got: %s", err)
	}
}

func TestReadAndSave(t *testing.T) {
	store := NewMemoryStorage()

	id := "id"
	doc := Document{Title: "some title"}
	store.Save(id, doc)

	v, err := store.Read(id)

	if v.Title != "some title" {
		t.Error("Expected document with correct title, got: %s", v.Title)
	}

	if err != nil {
		t.Error("Expected no error, got: %s", err)
	}
}

func TestReadAndSaveTwice(t *testing.T) {
	store := NewMemoryStorage()

	id := "id"
	doc := Document{Title: "some title"}
	store.Save(id, doc)

	doc2 := Document{Title: "some title2"}
	store.Save(id, doc2)

	v, err := store.Read(id)

	if v.Title != "some title2" {
		t.Error("Expected document with updated title, got: %s", v.Title)
	}

	if err != nil {
		t.Error("Expected no error, got: %s", err)
	}
}
