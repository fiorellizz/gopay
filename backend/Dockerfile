# ---------- STAGE 1: BUILD ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copia dependências
COPY go.mod ./
RUN go mod tidy

# Copia o resto do código
COPY . .

# Compila API
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api


# ---------- STAGE 2: RUNTIME ----------
FROM alpine:latest

WORKDIR /app

# Copia binário compilado
COPY --from=builder /app/api .

# Expõe porta
EXPOSE 8080

CMD ["./api"]