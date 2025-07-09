# Rinha Go API (Gin + Hexagonal)

## Rodando localmente

```bash
go run ./cmd/api
```

## Rodando com Docker

```bash
docker build -t rinha-go .
docker run -p 8080:8080 rinha-go
```

## Endpoints

- `GET /ping` — healthcheck
- `GET /business` — exemplo de uso da arquitetura hexagonal 