package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
)

type AesGcmEncryptor struct {
	key []byte
}

func NewAesGcmEncryptor(key []byte) *AesGcmEncryptor {
	return &AesGcmEncryptor{key: key}
}

func (e *AesGcmEncryptor) Encrypt(data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	encrypted := aesGCM.Seal(nonce, nonce, jsonData, nil)
	return encrypted, nil
}

func (e *AesGcmEncryptor) Decrypt(encryptedData []byte) (map[string]interface{}, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, errors.New("invalid encrypted data")
	}
	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	jsonData, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}
	return data, nil
}