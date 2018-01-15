package usecases

type ArticleNotFoundError error

type Export struct {
	Encoder       Encoder
	StorageClient Storer
}

func (e Export) ById(id string) (string, error) {
	return "", nil
}
