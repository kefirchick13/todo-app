FROM golang:latest

WORKDIR /app
RUN chmod -R 777 /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

ENV GOOS=linux
RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]