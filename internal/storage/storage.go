package storage

type IStorage interface {
	AddURL(url string) (string, error)
	GetURL(short string) (string, error)
}
