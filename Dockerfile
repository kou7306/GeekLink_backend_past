FROM golang:1.22.1

WORKDIR /app

COPY . /app

RUN go mod download

CMD ["go", "run", "main.go"]