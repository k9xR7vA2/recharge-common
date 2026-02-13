package constant

type DeviceType int

const (
	AndroidDevice DeviceType = iota + 1 //安卓
	IOSDevice                           //iOS
	BothDevice                          //双端
)

func (d DeviceType) String() string {
	switch d {
	case AndroidDevice:
		return "安卓"
	case IOSDevice:
		return "iOS"
	case BothDevice:
		return "双端"
	default:
		return "未知设备"
	}
}

func (d DeviceType) Code() string {
	switch d {
	case AndroidDevice:
		return "android"
	case IOSDevice:
		return "ios"
	case BothDevice:
		return "both"
	default:
		return "unknown"
	}
}

func (d DeviceType) IsValid() bool {
	return d >= AndroidDevice && d <= BothDevice
}
