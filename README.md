# go_gallery

Website for viewing photographs

## Setup

### Install Go

Follow the instructions on the [official website](https://golang.org/doc/install)

Add following path to your **.profile** 

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export PATH=$PATH:$(go env GOPATH)/bin
```

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

### Install Node.js

Node.js is required to run tailwindcss. Follow the instructions on the [Node.js official website](https://nodejs.org/en/download/) to install it.

To run tailwindcss in watch mode

```bash
npm run watch
```
