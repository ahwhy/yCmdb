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
go build -a -o dist/yCmdb -ldflags "-s -w" -ldflags "-X 'github.com/ahwhy/yCmdb/version.GIT_BRANCH=master' -X 'github.com/ahwhy/yCmdb/version.GIT_COMMIT=e7b449efc43787cb6be0fdbfe28f492487684772' -X 'github.com/ahwhy/yCmdb/version.BUILD_TIME=2021-11-29' -X 'github.com/ahwhy/yCmdb/version.GO_VERSION=v1.16.4'" main.go
./dist/yCmdb -f etc/cmdb-api.toml start
*/
