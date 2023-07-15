package storage

import (
	"fmt"
	"github.com/CvitoyBamp/urlshorter/internal/shortner"
	"log"
	"sync"
)

const (
	// Длина укороченного URL
	shortLen = 5
)

type Storage struct {
	sync.RWMutex
	store map[string]string
}

func CreateStorage() *Storage {
	return &Storage{
		store: make(map[string]string),
	}
}

func (s *Storage) AddURL(url string) (string, error) {

	short := shortner.RandUrlName(shortLen)

	s.RLock()
	for k, v := range s.store {

		if url == v {
			return "", fmt.Errorf("url was already added, use such url: %s", k)
		}

		if k == short {
			_, err := s.AddURL(url)
			if err != nil {
				return "", fmt.Errorf("can't add url")
			}
		}
	}
	s.RUnlock()

	s.Lock()
	s.store[short] = url
	s.Unlock()
	log.Printf("shorter %s for address %s was added to storage", short, url)

	return short, nil
}

func (s *Storage) GetURL(short string) (string, error) {
	s.RLock()
	defer s.RUnlock()

	val, exist := s.store[short]

	if exist {
		return val, nil
	}

	return "", fmt.Errorf("can't resolve to url")

}
