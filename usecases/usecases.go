package usecases

import "github.com/co0p/construe/storage"

type Decoder interface{}

type Encoder interface{}

type Translator interface {
	Translate(string) (string, error)
}

type Storer interface {
	Read(string) (storage.Document, error)
	Save(string, storage.Document) error
}

type Import struct {
	Decoder           Decoder
	TranslationClient Translator
	StorageClient     Storer
}
