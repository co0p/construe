package usecases

type Decoder interface{}

type Encoder interface{}

type Translator interface {
	Translate(string) (string, error)
}

type Storer interface {
	Read(string) ([]byte, error)
	Save(string, []byte) error
}

type Import struct {
	Decoder           Decoder
	TranslationClient Translator
	StorageClient     Storer
}
