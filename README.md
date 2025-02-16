# Encrypted Data Management Microservice

A secure, scalable, and extensible microservice with REST API for managing encrypted data, supporting multiple key providers and storage backends.

## Overview

This microservice provides:
- Data validation, encryption, storage, and decryption
- Integration with various cryptographic key providers (AWS KMS, HashiCorp Vault, 1Password)
- Support for multiple data stores (PostgreSQL, MongoDB, DynamoDB)
- Authentication, monitoring, logging, and versioning mechanisms

## Key Features

- **Secure Encryption**: AES-256-GCM encryption with physical separation of encrypted data
- **Multiple Key Providers**: Support for cloud and local cryptographic key sources
- **Flexible Storage**: Configurable data storage backends
- **Comprehensive Monitoring**: Authentication, structured logging, Prometheus metrics
- **Version Control**: Schema versioning and zero-downtime key rotation

## API Endpoints

### Store Data
```http
POST /store
```

Request:
```json
{
  "data_id": "optional",
  "schema": "user_schema_v2",
  "data": {
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

Response:
```json
{
  "data_id": "generated-or-specified-id"
}
```

### Retrieve Data
```http
POST /retrieve
```

Request:
```json
{
  "data_id": "user-123",
  "key_version": "optional"
}
```

Response:
```json
{
  "data": {
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

### Validate Data
```http
POST /validate
```

Request:
```json
{
  "schema": "user_schema_v2",
  "schema_version": 2,
  "data": {
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

Error Response:
```json
{
  "error": "invalid_schema",
  "details": ["Field 'email' is required"]
}
```

### Additional Endpoints

- `POST /batch/store`: Batch encryption and storage
- `POST /batch/retrieve`: Batch retrieval and decryption
- `GET /health`: Service health check
- `GET /ready`: Service readiness check

## Technical Stack

- **Language**: Go 1.20+
- **Framework**: Echo (with OpenAPI support)
- **Databases**: 
  - PostgreSQL (GORM)
  - MongoDB (official driver)
  - DynamoDB (AWS SDK)
  - SQLite (testing)
- **Configuration**: Viper + YAML
- **Containerization**: Docker, Kubernetes (Helm charts)
- **Monitoring**: Prometheus
- **Authentication**: JWT/OAuth2 with JWKS validation
- **Logging**: Structured JSON logging (zerolog)
- **Testing**: Go Test, testify, GoMock, chaos-mesh

## Configuration Examples

### AWS KMS with PostgreSQL

```yaml
security:
  auth:
    type: jwt
    jwks_url: "https://auth.example.com/.well-known/jwks.json"

key_provider:
  type: aws_kms
  aws_kms:
    region: "us-east-1"
    kms_key_id: "arn:aws:kms:us-east-1:123456789012:key/1234abcd-12ab-34cd-56ef-1234567890ab"

database:
  type: postgres
  migrations_path: "./migrations"
  url: "postgres://user:${VAULT:DB_PASSWORD}@db:5432/encrypted_data"

logging:
  level: "debug"
  format: "json"

monitoring:
  prometheus:
    enabled: true
    port: 9090
```

### HashiCorp Vault with MongoDB

```yaml
security:
  auth:
    type: jwt
    jwks_url: "https://auth.example.com/.well-known/jwks.json"

key_provider:
  type: vault
  vault:
    address: "https://vault.example.com"
    transit_key: "encryption_key_v2"

database:
  type: mongodb
  uri: "mongodb://db_user:db_password@db.example.com:27017"
  database: "encrypted_data_db"

logging:
  level: "info"
  format: "json"

monitoring:
  prometheus:
    enabled: true
    port: 9090
```

## Architecture

The service follows Domain-Driven Design principles with a layered architecture:

- **Domain Layer**: Core entities, value objects, and interfaces
- **Application Layer**: Use cases for encryption, decryption, and validation
- **Infrastructure Layer**: Key provider implementations, database adapters
- **Interface Layer**: REST API, CLI tools, optional gRPC interface

## Non-Functional Requirements

### Security
- TLS 1.3 encryption
- Secrets stored in HashiCorp Vault
- No key logging or client exposure
- Rate limiting (100 RPM per IP/token)

### Reliability & Scalability
- Automatic retry for key provider/database failures
- Horizontal scaling support
- Health check endpoints

### Performance
- 1000+ RPS per instance (c4.large)
- Optimized cryptographic operations

## Development Timeline

1. **Core & Domain** (2 weeks)
   - Domain model design
   - Basic encryption/decryption logic

2. **Application Layer & API** (2 weeks)
   - Core operations implementation
   - REST API development
   - Authentication integration

3. **Infrastructure** (2 weeks)
   - Key provider adapters
   - Database implementations
   - Configuration system

4. **Security & Monitoring** (1 week)
   - Logging and metrics
   - Security features
   - Container setup

5. **Testing** (2 weeks)
   - Unit tests (80%+ coverage)
   - Integration tests
   - Load and chaos testing

## Acceptance Criteria

- 1000+ RPS per instance
- Zero key leaks in logs/responses
- Hot-reload configuration support
- Complete API and configuration documentation
- Monitoring system integration

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the [LICENSE.md](LICENSE.md) file for details.
