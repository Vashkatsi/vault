package config

import "os"

type Config struct {
	RepositoryType   string
	DBUrl            string
	EncryptorType    string
	EncryptionKey    string
	StoreEndpoint    string
	RetrieveEndpoint string
	Port             string
}

func LoadConfig() Config {
	return Config{
		RepositoryType:   getEnv("REPOSITORY_TYPE", "in_memory"),
		DBUrl:            os.Getenv("DB_URL"),
		EncryptorType:    getEnv("ENCRYPTOR_TYPE", "dummy"),
		EncryptionKey:    os.Getenv("ENCRYPTION_KEY"),
		StoreEndpoint:    getEnv("STORE_ENDPOINT", "/store"),
		RetrieveEndpoint: getEnv("RETRIEVE_ENDPOINT", "/retrieve"),
		Port:             getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}
