package storage

import (
	"errors"
	"sync"

	"github.com/Vashkatsi/vault/internal/domain"
)

type InMemoryRepository struct {
	data map[string]*domain.EncryptedData
	mu   sync.Mutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: make(map[string]*domain.EncryptedData),
	}
}

func (repo *InMemoryRepository) Save(data *domain.EncryptedData) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.data[data.DataID] = data
	return nil
}

func (repo *InMemoryRepository) Retrieve(dataID string) (*domain.EncryptedData, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if data, ok := repo.data[dataID]; ok {
		return data, nil
	}
	return nil, errors.New("data not found")
}