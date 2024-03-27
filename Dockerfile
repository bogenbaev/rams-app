FROM golang:latest

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go get -u github.com/swaggo/swag/cmd/swag

RUN go build -o main

# ENV MY_ENV_VAR=value

EXPOSE 8080


CMD ["./main"]
