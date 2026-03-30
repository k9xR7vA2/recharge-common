package schema

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/schema/business"
	"github.com/k9xR7vA2/recharge-common/schema/types"
	"sync"
)

var (
	registry = map[constant.BusinessType][]types.RawField{} // ← 改
	once     sync.Once
)

func initRegistry() {
	once.Do(func() {
		for _, bs := range []types.BusinessSchema{
			business.Mobile,
			business.IndiaMobile,
			business.IndiaDTH,
		} {
			registry[bs.BusinessType] = bs.Fields
		}
	})
}
