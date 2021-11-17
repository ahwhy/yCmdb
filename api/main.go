package main

import (
	"github.com/ahwhy/yCmdb/api/cmd"
)

func main() {
	cmd.Execute()
}

/*
go run main.go -f etc/demo.toml start
go build -ldflags "-s -w" -o yCmdb main.go
go build -o yCmdb -ldflags "-X github.com/ahwhy/yCmdb/api/version.GIT_TAG='v0.1.1'" main.go
*/
