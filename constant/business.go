package constant

// BusinessType 业务类型
type BusinessType string

const (
	MobileBusiness   BusinessType = "mobile"       // MobileBusiness 话费
	ElectricBusiness BusinessType = "electric"     // ElectricBusiness 电费
	GameBusiness     BusinessType = "game"         // GameBusiness 游戏
	CardBusiness     BusinessType = "card"         // CardBusiness 卡密
	OilBusiness      BusinessType = "oil"          // 加油业务
	IndiaMobile      BusinessType = "india_mobile" // 印度话费
	IndiaDTH         BusinessType = "india_dth"    // 印度 DTH (卫星电视) 充值
)

func (b BusinessType) String() string {
	switch b {
	case MobileBusiness:
		return "mobile"
	case ElectricBusiness:
		return "electric"
	case GameBusiness:
		return "game"
	case CardBusiness:
		return "card"
	case OilBusiness:
		return "oil"
	case IndiaMobile:
		return "india_mobile"
	case IndiaDTH:
		return "india_dth"
	default:
		return "未知业务"
	}
}

// ShowName   实现匹配规则的字符串表示
func (b BusinessType) ShowName() string {
	switch b {
	case MobileBusiness:
		return "话费"
	case ElectricBusiness:
		return "电费"
	case GameBusiness:
		return "游戏"
	case CardBusiness:
		return "卡密"
	case OilBusiness:
		return "加油"
	case IndiaMobile:
		return "印度话费"
	case IndiaDTH:
		return "印度DTH"
	default:
		return "未知业务"
	}
}

// IsValid 验证匹配规则是否有效
func (b BusinessType) IsValid() bool {
	switch b {
	case MobileBusiness, ElectricBusiness, GameBusiness, OilBusiness, CardBusiness, IndiaMobile, IndiaDTH:
		return true
	default:
		return false
	}
}

// GetAllBusinessTypes 获取所有业务类型
func GetAllBusinessTypes() []struct {
	Label string       `json:"label"`
	Value BusinessType `json:"value"`
} {
	return []struct {
		Label string       `json:"label"`
		Value BusinessType `json:"value"`
	}{
		{Label: MobileBusiness.ShowName(), Value: MobileBusiness},
		{Label: ElectricBusiness.ShowName(), Value: ElectricBusiness},
		{Label: GameBusiness.ShowName(), Value: GameBusiness},
		{Label: CardBusiness.ShowName(), Value: CardBusiness},
		{Label: OilBusiness.ShowName(), Value: OilBusiness},
		{Label: IndiaMobile.ShowName(), Value: IndiaMobile},
		{Label: IndiaDTH.ShowName(), Value: IndiaDTH},
	}
}
