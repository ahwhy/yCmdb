package resource

import "context"

type Service interface {
	// 搜索Resource
	Search(context.Context, *SearchRequest) (*ResourceSet, error)
}
