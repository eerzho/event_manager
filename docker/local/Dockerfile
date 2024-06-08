FROM golang:1.22

WORKDIR /app
COPY . .

RUN go mod download

CMD ["sh", "-c", "while :; do sleep 1; done"]
