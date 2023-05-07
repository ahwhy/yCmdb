package client_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ahwhy/yCmdb/apps/resource"
	"github.com/ahwhy/yCmdb/client"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)
	conf := client.NewConfig("localhost:18060")
	conf.WithClientCredentials("nHerVBlrKIDurviMGUXVOQHC", "l5FB38Mw2JmxHgGm8rUcich2ZrGRVrl7")

	c, err := client.NewClient(conf)
	if should.NoError(err) {
		rs, err := c.Resource().Search(context.Background(), resource.NewSearchRequest())
		should.NoError(err)
		fmt.Println(rs)
	}
}
