package constant

type IndiaDthOperatorType int

const (
	DthTataPlay  IndiaDthOperatorType = 106
	DthAirtel    IndiaDthOperatorType = 101
	DthDishTV    IndiaDthOperatorType = 105
	DthSunDirect IndiaDthOperatorType = 103
)

func (d IndiaDthOperatorType) IsValid() bool {
	switch d {
	case DthTataPlay, DthAirtel, DthDishTV, DthSunDirect:
		return true
	}
	return false
}

func (d IndiaDthOperatorType) String() string {
	switch d {
	case DthTataPlay:
		return "Tata Play"
	case DthAirtel:
		return "Airtel Digital TV"
	case DthDishTV:
		return "Dish TV"
	case DthSunDirect:
		return "Sun Direct"
	default:
		return "未知"
	}
}

func (d IndiaDthOperatorType) Code() string {
	switch d {
	case DthTataPlay:
		return "tata_play"
	case DthAirtel:
		return "airtel"
	case DthDishTV:
		return "dish_tv"
	case DthSunDirect:
		return "sun_direct"
	default:
		return "unknown"
	}
}
