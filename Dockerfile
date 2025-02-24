FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/ cmd/
COPY internal/ internal/
RUN go mod download
RUN go build -o vault ./cmd/vault/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/vault .
EXPOSE 8080
ENTRYPOINT ["./vault"]