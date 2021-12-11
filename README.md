# My_Cmdb

## 多云资产管理平台, 支持厂商
- 阿里云
- 腾讯云
- 华为云
- VMware

## Grpc-Client
```go
package main

import (
	"context"
	"fmt"

	"github.com/ahwhy/yCmdb/client"
	"github.com/ahwhy/yCmdb/app/resource"
)

func main() {
    // 配置yCmdb grpc服务调用地址和凭证
	conf := client.NewConfig("localhost:18060")
	conf.WithClientCredentials("xx", "xx")

    // 创建yCmdb客户端
	yCmdb, err := client.NewClient(conf)
	if err != nil {
		panic(err)
	}

    // 服务调用
	rs, err := yCmdb.Resource().Search(context.Background(), resource.NewSearchRequest())
	if err != nil {
		panic(err)
	}
	fmt.Println(rs)
}
```