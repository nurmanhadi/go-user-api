FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod download
RUN go build -o ./build/app ./cmd/web

EXPOSE 3000
CMD [ "./build/app" ]