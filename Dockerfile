# --- Stage 1: Builder на основе Alpine + Golang
FROM golang:1.23.11-alpine AS builder

WORKDIR /src

# Кэширование зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Сборка бинарника (без CGO, для минимального образа)
RUN CGO_ENABLED=0 GOOS=linux go build -o startbase ./cmd/server

# --- Stage 2: Production образ
FROM alpine:latest

# Устанавливаем корневые сертификаты
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем только нужные артефакты из builder
COPY --from=builder /src/startbase .
COPY --from=builder /src/templates ./templates

# 🔒 Открываем порт приложения
EXPOSE 8080

# ✅ Продакшн-режим для Gin (по желанию)
ENV GIN_MODE=release

# 🚦 Добавляем простой healthcheck (опционально)
HEALTHCHECK --interval=30s --timeout=5s \
  CMD wget --spider http://localhost:8080/health || exit 1

# 🔁 Запуск
ENTRYPOINT ["./startbase"]