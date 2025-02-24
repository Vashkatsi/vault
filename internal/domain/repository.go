package domain

type Repository interface {
	Save(data *EncryptedData) error
	Retrieve(dataID string) (*EncryptedData, error)
}