.PHONY: mod-tidy build run docker-build docker-run docker-compose-up

mod-tidy:
	go mod tidy

build: mod-tidy
	go build -o vault ./cmd/vault/main.go

run: build
	./vault

docker-build:
	docker build -t vault .

docker-run:
	docker run --rm -p 8080:8080 vault

docker-compose-up:
	cd deployments/docker && docker-compose up --build