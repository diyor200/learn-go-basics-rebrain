FROM golang:1.20-alpine

WORKDIR /app

COPY main.go .

CMD [ "go", "run", "main.go" ]