package schema

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/schema/business"
	"sync"
)

// BusinessSchema 一个业务类型的完整定义
type BusinessSchema struct {
	BusinessType constant.BusinessType
	Fields       []RawField
}

var (
	registry = map[constant.BusinessType][]RawField{}
	once     sync.Once
)

func initRegistry() {
	once.Do(func() {
		for _, bs := range []BusinessSchema{
			business.Mobile,
			business.IndiaMobile,
			business.IndiaDTH,
			// 新增业务类型在这里加一行
		} {
			registry[bs.BusinessType] = bs.Fields
		}
	})
}
