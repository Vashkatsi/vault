package initdeps

import (
	"encoding/hex"
	"log"
	"os"

	"github.com/Vashkatsi/vault/internal/application"
	"github.com/Vashkatsi/vault/internal/domain"
	"github.com/Vashkatsi/vault/internal/infrastructure/encryption"
	"github.com/Vashkatsi/vault/internal/infrastructure/storage"
)

func InitializeDependencies() *application.DataService {
	repo := initializeRepository()
	encryptor := initializeEncryptor()
	return application.NewDataService(repo, encryptor)
}

func initializeRepository() domain.Repository {
	repoType := os.Getenv("REPOSITORY_TYPE")
	if repoType == "postgres" {
		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			log.Fatal("DB_URL must be provided for postgres repository")
		}
		postgresRepo, err := storage.NewPostgresRepository(dbURL)
		if err != nil {
			log.Fatalf("Failed to initialize Postgres repository: %v", err)
		}
		return postgresRepo
	}

	return storage.NewInMemoryRepository()
}

func initializeEncryptor() encryption.Encryptor {
	encryptorType := os.Getenv("ENCRYPTOR_TYPE")
	if encryptorType == "aes_gcm" {
		keyHex := os.Getenv("ENCRYPTION_KEY")
		if keyHex == "" {
			log.Fatal("ENCRYPTION_KEY must be provided for aes_gcm encryptor")
		}
		key, err := hex.DecodeString(keyHex)
		if err != nil {
			log.Fatalf("Invalid ENCRYPTION_KEY: %v", err)
		}
		return encryption.NewAesGcmEncryptor(key)
	}
	
	return encryption.NewDummyEncryptor()
}