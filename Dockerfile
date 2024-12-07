FROM golang:alpine

# Убедимся, что у нас есть необходимые инструменты
RUN apk update && apk add --no-cache \
    bash \
    curl \
    postgresql-client \
    build-base \
    git

# Установить версию Go (дополнительно, если нужно)
RUN go version

# Установить GOPATH (опционально)
ENV GOPATH=/

# Копируем исходный код приложения
COPY ./ ./

# Делаем скрипт для ожидания Postgres исполняемым
RUN chmod +x wait-for-postgres.sh

# Загружаем зависимости Go
RUN go mod download

# Сборка Go-приложения
RUN go build -o todo-app ./cmd/main.go

# Указываем команду запуска приложения
CMD ["./todo-app"]


# # Этап сборки
# FROM golang:1.22.4 AS builder

# WORKDIR /app

# # Копируем зависимости
# COPY go.mod go.sum ./
# RUN go mod download

# # Копируем исходный код и собираем приложение
# COPY . .
# RUN go build -o go-todo-app .

# # Этап выполнения
# FROM alpine:latest

# COPY --from=builder /app/go-todo-app /go-todo-app

# RUN chmod +x /go-todo-app

# CMD ["/go-todo-app"]