package application

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/Vashkatsi/vault/internal/domain"
	"github.com/Vashkatsi/vault/internal/infrastructure/encryption"
)

type DataService struct {
	repository domain.Repository
	encryptor  encryption.Encryptor
}

func NewDataService(repo domain.Repository, encryptor encryption.Encryptor) *DataService {
	return &DataService{
		repository: repo,
		encryptor:  encryptor,
	}
}

func (s *DataService) StoreData(dataID string, plainData map[string]interface{}) (string, error) {
	if dataID == "" {
		newID, err := generateID()
		if err != nil {
			return "", err
		}
		dataID = newID
	}
	encryptedBytes, err := s.encryptor.Encrypt(plainData)
	if err != nil {
		return "", err
	}
	ed := &domain.EncryptedData{
		DataID:           dataID,
		EncryptedContent: encryptedBytes,
	}
	if err := s.repository.Save(ed); err != nil {
		return "", err
	}
	return dataID, nil
}

func (s *DataService) RetrieveData(dataID string) (map[string]interface{}, error) {
	ed, err := s.repository.Retrieve(dataID)
	if err != nil {
		return nil, err
	}
	plainData, err := s.encryptor.Decrypt(ed.EncryptedContent)
	if err != nil {
		return nil, err
	}
	return plainData, nil
}

func generateID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}