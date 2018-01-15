package storage

type MemoryStorage struct{}

func (m MemoryStorage) Read(key string) ([]byte, error) {
	return []byte{0}, nil
}

func (m MemoryStorage) Save(key string, data []byte) error {
	return nil
}
