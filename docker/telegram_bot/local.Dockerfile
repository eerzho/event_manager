FROM golang:1.22

WORKDIR /telegram_bot
COPY . .

RUN go mod download

CMD ["sh", "-c", "while :; do sleep 1; done"]
