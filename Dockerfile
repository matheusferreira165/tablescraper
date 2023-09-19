FROM golang:alpine

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd

ENTRYPOINT [ "/app/server" ]