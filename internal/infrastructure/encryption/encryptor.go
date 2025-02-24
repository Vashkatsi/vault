package encryption

type Encryptor interface {
	Encrypt(data map[string]interface{}) ([]byte, error)
	Decrypt(encryptedData []byte) (map[string]interface{}, error)
}