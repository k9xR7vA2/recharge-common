package constant

type IndiaCarrierType int //运营商类型1联通，2移动，3电信

const (
	Airtel IndiaCarrierType = iota + 1 //中国联通
	Jio                                // 中国移动
)

// String 实现 Stringer 接口，返回运营商名称的字符串描述
func (s IndiaCarrierType) String() string {
	switch s {
	case Airtel:
		return "Airtel"
	case Jio:
		return "Jio"
	default:
		return "Unknown"
	}
}

func (s IndiaCarrierType) Int() int {
	switch s {
	case Airtel:
		return 1
	case Jio:
		return 2
	default:
		return 0
	}
}

// Code  实现 Stringer 接口，返回运营商编码的字符串描述
func (s IndiaCarrierType) Code() string {
	switch s {
	case Airtel:
		return "airtel"
	case Jio:
		return "jio"
	default:
		return "Unknown"
	}
}

func (s IndiaCarrierType) IsValid() bool {
	return s >= Airtel && s <= Jio
}
