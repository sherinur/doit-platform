package usecase

type FileRepo interface {
	Create(object []byte) (string, error)
	Get(key string) ([]byte, error)
	Delete(key string) error
}
