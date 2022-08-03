# Technical challenge for Stori

See [Instructions of the challenge](Challenge.md)

## Requirements

- Go 1.18
- Docker
- `docker compose` (v2)

## About this project

This project has a [docker-compose.yml](docker-compose.yml) that creates 3 containers:

1. **mongodb:** MongoDB instance running and initialize a database and a collection with the values at `.env` file
2. **mongo-express:** a lightweight web-based administrative interface deployed to manage MongoDB databases interactively at `http://localhost:8081`
3. **go-app:** a Go app that:
    - reads the transactions in the [txns.csv](txns.csv) file
    - inserts theses transactions into the MongoDB container
    - processes and makes the calculations specified at [Challenge.md](Challenge.md) and outputs a summary information
    - creates `email.html` containing the summary information

## How to execute it

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

Expected output:

```bash
...
Attaching to go-app
go-app  | Transaction added: &{ObjectID("62e96aa00ba6ff6fdeafb015")}
go-app  | Transaction added: &{ObjectID("62e96aa00ba6ff6fdeafb016")}
go-app  | Transaction added: &{ObjectID("62e96aa00ba6ff6fdeafb017")}
go-app  | Transaction added: &{ObjectID("62e96aa00ba6ff6fdeafb018")}
go-app  | Total balance is: 39.74
go-app  | Number of transactions in July: 2
go-app  | Number of transactions in August: 2
go-app  | Average debit amount: -15.38
go-app  | Average credit amount: 35.25
go-app  | Email template created: /template/email.html
go-app exited with code 0
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

## How to upload to AWS Lambda

1. Build the code for lambda

```bash
make build
```

2. Upload `bin/main.zip` to AWS Lambda

The `txns.csv` and `email-template.html` are stored as public at AWS S3:

- [https://s3.amazonaws.com/rodomar.mx/assets/files/txns.csv](https://s3.amazonaws.com/rodomar.mx/assets/files/txns.csv)
- [https://s3.amazonaws.com/rodomar.mx/assets/files/email-template.html](https://s3.amazonaws.com/rodomar.mx/assets/files/email-template.html)
