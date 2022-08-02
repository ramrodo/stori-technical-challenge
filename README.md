# Technical challenge for Stori

See [Instructions of the challenge](Challenge.md)

## Requirements

- Go 1.18
- Docker
- `docker compose` (v2)

## How to execute it

### Local

```bash
go run main.go
```

### Docker

1. Build containers

```bash
docker compose build
```

2. Run DB containers (`mongo` and `mongo-express`)

```bash
docker compose up mongodb mongo-express
```

3. Wait until MongoDB starts and initializes and then, in another terminal, run Go app:

```bash
docker compose up go-app
```

#### To stop and remove Docker stuff created

1. Remove containers

```bash
docker compose down
```

2. Remove volumes

```bash
docker volume ls
docker volume rm <volume-name>
```
