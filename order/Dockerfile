FROM golang:1.16

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /server cmd/api/main.go

ENTRYPOINT ["/server"]