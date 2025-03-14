FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./build/app ./cmd/web/main.go

EXPOSE 3000

CMD ./build/app