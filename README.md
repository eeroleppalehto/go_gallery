# go_gallery

Webapp for viewing photographs

## Setup

To try this app locally create a `.env` file in the root directory with the following content

```bash
MYSQL_ROOT_PASSWORD="password"
DATABASE_URL="root:password@tcp(gollery_db:3306)/gollery?parseTime=true"
```

After creating the `.env` file, run the following command to start the server

```bash
docker compose up -d
```

## Development

To develop this app, you need to install the following tools

### Install Go

Follow the instructions on the [official website](https://golang.org/doc/install) to install Go.

### Install Air

Air is a live reload tool for Go applications. It is used to automatically reload the server when changes are made to the code. Install it using the following command

```bash
go install github.com/cosmtrek/air@latest
```

Run the following command to start the server

```bash
air
```

### Install templ

templ is a tool for generating Go templates. Install it using the following command

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Install sqlc

**sqlc** is a tool for generating Go code from SQL queries. Install it using the following command

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### Install Node.js

Node.js is required to run tailwindcss. Follow the instructions on the [Node.js official website](https://nodejs.org/en/download/) to install it.

To run tailwindcss in watch mode

```bash
npm run watch
```

### Setup MySQL Server

To setup a MySQL server using Docker, run the following commands

```bash
docker build --tag gollery_db ./database
```

```bash
docker run --name gollery_db -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 gollery_db -d
```

### Setup environment variables

Add the following environment variables to a `.env` file in the root directory

```bash
DATABASE_URL_DEV="root:Q2werty@tcp(127.0.0.1:3306)/gollery?parseTime=true"
ENV="DEV"
```
