package storage

import (
	"github.com/Vashkatsi/vault/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(dbURL string) (*PostgresRepository, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.EncryptedData{}); err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Save(data *domain.EncryptedData) error {
	return r.db.Create(data).Error
}

func (r *PostgresRepository) Retrieve(dataID string) (*domain.EncryptedData, error) {
	var data domain.EncryptedData
	err := r.db.First(&data, "data_id = ?", dataID).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}