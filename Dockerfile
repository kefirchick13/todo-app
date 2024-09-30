FROM golang:1.22.5-alpine AS compiler
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /build/app 

FROM alpine
WORKDIR /app
COPY --from=compiler ./build/app /app/app

RUN chmod -R 777 /app/app
CMD [ "/app/app" ]

# FROM golang:latest 

# COPY ./ ./
# RUN go build -o main

# RUN chmod +x ./main

# CMD [ "./main" ]



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