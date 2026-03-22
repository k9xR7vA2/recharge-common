package types

// DictOption 前端字典选项标准格式
type DictOption struct {
	Label        string            `json:"label"`                // 显示文本
	Value        interface{}       `json:"value"`                // 选项值
	Code         string            `json:"code"`                 // 编码（可选）
	RechargeMode string            `json:"recharge_mode"`        // 新增
	ExtParams    map[string]string `json:"ext_params,omitempty"` // 新增
}

// DictResponse 字典响应标准格式
type DictResponse struct {
	Key     string       `json:"key"`     // 字典键名
	Name    string       `json:"name"`    // 字典名称
	Options []DictOption `json:"options"` // 选项列表
}
