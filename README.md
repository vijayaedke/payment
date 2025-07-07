# Payment App

# Customer Account & Transactions

A customer has an account. For each operation performed by the customer, a
transaction is created and associated with their respective account.
Transactions are specified by an identifier, a type, an amount, and a creation date.
The available types are purchase, installment purchase, withdrawal, and payments.
Purchase, installment purchase, and withdrawal transaction types are stored with
negative amounts (debt transactions), while payments are stored with positive
amounts (credit transactions).

## APIs

- Implement a REST endpoint for creating accounts - POST /accounts.
- Implement a REST endpoint for retrieving existing accounts - GET /accounts/:accountId.
- Implement a REST endpoint for creating transactions - POST /transactions.

## Prerequisites

Before running the application, ensure you have the following installed:

- **Go**: Version 1.22 or higher (check with `go version`).
- **Git**: For cloning or interacting with the Go modules.

## Setup

## 1. Clone the repository:

Clone this project to your local machine.

```bash
git clone https://github.com/vijayaedke/payment
cd payment
```

## 2. Install dependencies

This project uses Go modules. Ensure the `go.mod` and `go.sum` files are properly set up (these files should already be present).

Install the required Go modules using:

```bash
go mod tidy
```

## 3. Run the application

Once the dependencies are installed and Go modules are initialized, run the mysql DB config using `mysql_dump.sql` file commands and you can run the Go program.

```bash
#if mysql is not available locally, pull the latest or 8.0 and above version and run it using following creds
docker pull mysql:8.0

docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=payment -e MYSQL_USER=root -e MYSQL_PASSWORD=root -p 3306:3306 -d mysql:8.0

go run main.go
```

## 4. Swagger

View API swagger documentation using swagger editor https://editor.swagger.io/
Copy paste the contents of `swagger.yaml` file into the above editor.

## 5. Run using docker compose

The below commands will remove any existing build to be discarded and build a new app. Run the docker image using the run command and
access the APIs using `http://localhost:9090` as the base the url.

```bash
    # remove any unwanted docker-compose build running
    docker-compose down -v
    docker system prune -f

    docker-compose up --build -d

    docker-compose run -d -p 9090:9090 app ./main payment
```
