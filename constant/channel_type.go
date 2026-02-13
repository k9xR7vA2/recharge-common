package constant

type ChannelType int

const (
	ChannelTypeBasic     ChannelType = iota + 1 //基础通道
	ChannelTypeComposite                        //混合通道
)

func (d ChannelType) String() string {
	switch d {
	case ChannelTypeBasic:
		return "基础通道"
	case ChannelTypeComposite:
		return "混合通道"
	default:
		return "未知类型"
	}
}

func (d ChannelType) Val() int {
	switch d {
	case ChannelTypeBasic:
		return 1
	case ChannelTypeComposite:
		return 2
	default:
		return 0
	}
}

func (d ChannelType) Code() string {
	switch d {
	case ChannelTypeBasic:
		return "basic"
	case ChannelTypeComposite:
		return "composite"
	default:
		return "Unknown"
	}
}

func (d ChannelType) IsValid() bool {
	return d >= ChannelTypeBasic && d <= ChannelTypeComposite
}
