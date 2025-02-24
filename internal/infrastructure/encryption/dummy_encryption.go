package encryption

import (
	"encoding/json"
	"errors"
)

type DummyEncryptor struct{}

func NewDummyEncryptor() *DummyEncryptor {
	return &DummyEncryptor{}
}

func (e *DummyEncryptor) Encrypt(data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	
	return append([]byte("encrypted:"), jsonData...), nil
}

func (e *DummyEncryptor) Decrypt(encryptedData []byte) (map[string]interface{}, error) {
	prefix := []byte("encrypted:")
	if len(encryptedData) < len(prefix) || string(encryptedData[:len(prefix)]) != "encrypted:" {
		return nil, errors.New("invalid encrypted data")
	}
	jsonData := encryptedData[len(prefix):]
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}
	return data, nil
}