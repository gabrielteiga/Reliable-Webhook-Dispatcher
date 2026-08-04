# Reliable-Webhook-Dispatcher

Receives webhook requests and delivers them asynchronously and reliably.

## Prerequisites

- **Go 1.25**
- **air** — only for hot reload:

  ```bash
  go install github.com/air-verse/air@latest
  ```

- **swag** — only to generate the docs:

  ```bash
  go install github.com/swaggo/swag/cmd/swag@v1.16.4
  ```

## Running with air (hot reload)

```bash
air
```

## Running with build

```bash
make run
```

Or build only:

```bash
make build
```

The server starts at `http://localhost:8080`.

## Documentation (Swagger)

Generate:

```bash
make swagger
```

Access with the server running at:

```
http://localhost:8080/swagger/index.html
```
