# Project Setup and Launch Instructions

## Prerequisites
1. Create a `.env` file based on the example provided in `.env.example`.
2. Download and install [ngrok](https://ngrok.com/).
3. Run `ngrok http TELEGRAM_PORT` and add the resulting URL to your `.env` file under the key `DOMAIN`.

## Local Launch

1. Install dependencies:
    ```bash
    make deps
    ```

2. Run the server:
    ```bash
    make run
    ```

   The HTTP server will be running on the port specified by `HTTP_PORT`, and the Telegram bot will be running on the port specified by `TELEGRAM_PORT`.

## Launch with Docker Compose

1. Build and start the containers:
    ```bash
    docker compose -f docker-compose.app.local.yaml up --detach --build
    ```

2. Access the container's console:
    ```bash
    docker compose exec app sh
    ```

3. Run the application:
    ```bash
    go run ./cmd/app
    ```

   The HTTP server will be running on the port specified by `HTTP_PORT`, and the Telegram bot will be running on the port specified by `TELEGRAM_PORT`.
