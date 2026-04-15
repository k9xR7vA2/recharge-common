package types

type RuleConfig interface {
	Validate() error
	ExtractChannelID(amount string) (uint, error)
	ExtractAllChannelIDsSortedByWeight() ([]uint, error)
	GetSupportedAmounts() []int // 返回策略支持的所有金额
	GetAllChannelIDs() []int
}

// ChannelProcessor 定义通道处理器接口
type ChannelProcessor interface {
	ProcessAmount(config RuleConfig, baseChannelMap map[uint]ChannelAmount) []ChannelAmount
}
