package types

// ChannelAmount 通道金额配置
type ChannelAmount struct {
	ChannelName string             `json:"channel_name"` // 通道名称
	Options     []AmountOptionResp `json:"options"`      // 该通道的金额选项列表
}

type AmountOptionResp struct {
	Value int `json:"value"`
	Label int `json:"label"`
}
