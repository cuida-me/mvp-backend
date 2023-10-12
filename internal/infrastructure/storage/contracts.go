package storage

type Provider interface {
	Save(filename, folder string) error
	Delete(filename, folder string) error
}
