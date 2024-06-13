FROM golang:1.22

WORKDIR /http_server
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod download

CMD ["sh", "-c", "while :; do sleep 1; done"]
