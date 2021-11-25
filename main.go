package main

import (
	"github.com/ahwhy/yCmdb/cmd"
)

func main() {
	cmd.Execute()
}

/*
go run main.go -f etc/demo.toml start
go build -ldflags "-s -w" -o yCmdb main.go
go build -o ./bin/yCmdb -ldflags "-X github.com/ahwhy/yCmdb/version.GIT_TAG='v1.0.0'" main.go
./bin/yCmdb -f etc/cmdb-api.toml start
*/
