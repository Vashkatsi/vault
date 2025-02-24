package domain

type EncryptedData struct {
	DataID           string `gorm:"primaryKey"`
	EncryptedContent []byte
}