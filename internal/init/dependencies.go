package initdeps

import (
	"encoding/hex"
	"github.com/Vashkatsi/vault/internal/application"
	"github.com/Vashkatsi/vault/internal/config"
	"github.com/Vashkatsi/vault/internal/domain"
	"github.com/Vashkatsi/vault/internal/infrastructure/encryption"
	"github.com/Vashkatsi/vault/internal/infrastructure/storage"
	"log"
)

func InitializeDependencies() *application.DataService {
	cfg := config.LoadConfig()
	repo := initializeRepository(cfg)
	encryptor := initializeEncryptor(cfg)
	return application.NewDataService(repo, encryptor)
}

func initializeRepository(cfg *config.Config) domain.Repository {
	if cfg.RepositoryType == "postgres" {
		if cfg.DBUrl == "" {
			log.Fatal("DB_URL must be provided for postgres repository")
		}
		postgresRepo, err := storage.NewPostgresRepository(cfg.DBUrl)
		if err != nil {
			log.Fatalf("Failed to initialize Postgres repository: %v", err)
		}
		return postgresRepo
	}
	return storage.NewInMemoryRepository()
}

func initializeEncryptor(cfg *config.Config) encryption.Encryptor {
	if cfg.EncryptorType == "aes_gcm" {
		if cfg.EncryptionKey == "" {
			log.Fatal("ENCRYPTION_KEY must be provided for aes_gcm encryptor")
		}
		key, err := hex.DecodeString(cfg.EncryptionKey)
		if err != nil {
			log.Fatalf("Invalid ENCRYPTION_KEY: %v", err)
		}
		return encryption.NewAesGcmEncryptor(key)
	}
	return encryption.NewDummyEncryptor()
}
