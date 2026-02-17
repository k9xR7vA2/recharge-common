package constant

type IndiaCarrierType int

const (
	Airtel IndiaCarrierType = iota + 1
	Jio

	VI
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
	case VI:
		return "vi"
	default:
		return "Unknown"
	}
}

func (s IndiaCarrierType) IsValid() bool {
	return s >= Airtel && s <= VI
}

func IndiaCarrierTypeFromLookup(ct string) (IndiaCarrierType, bool) {
	switch ct {
	case "airtel":
		return Airtel, true
	case "jio":
		return Jio, true
	case "vi":
		return VI, true
	default:
		return 0, false
	}
}