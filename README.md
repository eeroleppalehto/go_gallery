# go_gallery
Website for viewing photographs

## Setup

Install air for live reload
```bash
go install github.com/cosmtrek/air@latest
```

Add the following to your .bashrc or .zshrc
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Run the following command to start the server
```bash
air
```

Also run tailwindcss in watch mode
```bash
npm run watch
```
