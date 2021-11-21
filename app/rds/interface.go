package rds

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewRdsSet(),
	}
}

type PagerResult struct {
	Data    *RdsSet
	Err     error
	HasNext bool
}
