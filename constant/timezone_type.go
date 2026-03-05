package constant

import "time"

// TimezoneType 时区类型
type TimezoneType int

const (
	TimezoneShanghai   TimezoneType = 1 + iota // 亚洲/上海 UTC+8       中国
	TimezoneKolkata                            // 亚洲/加尔各答 UTC+5:30  印度
	TimezoneNewYork                            // 美洲/纽约 UTC-5         美国东部
	TimezoneLosAngeles                         // 美洲/洛杉矶 UTC-8       美国西部
	TimezoneLondon                             // 欧洲/伦敦 UTC+0         英国
	TimezoneDubai                              // 亚洲/迪拜 UTC+4         迪拜
	TimezoneManila                             // 亚洲/马尼拉 UTC+8       菲律宾
)

// timezoneIANAMap ID → IANA时区映射
var timezoneIANAMap = map[TimezoneType]string{
	TimezoneShanghai:   "Asia/Shanghai",
	TimezoneKolkata:    "Asia/Kolkata",
	TimezoneNewYork:    "America/New_York",
	TimezoneLosAngeles: "America/Los_Angeles",
	TimezoneLondon:     "Europe/London",
	TimezoneDubai:      "Asia/Dubai",
	TimezoneManila:     "Asia/Manila",
}

// timezoneLabelMap ID → 显示名称映射
var timezoneLabelMap = map[TimezoneType]string{
	TimezoneShanghai:   "中国标准时间 (UTC+8)",
	TimezoneKolkata:    "印度标准时间 (UTC+5:30)",
	TimezoneNewYork:    "美东时间 (UTC-5)",
	TimezoneLosAngeles: "美西时间 (UTC-8)",
	TimezoneLondon:     "英国时间 (UTC+0)",
	TimezoneDubai:      "迪拜时间 (UTC+4)",
	TimezoneManila:     "菲律宾时间 (UTC+8)",
}

// ToIANA 转换为IANA标准格式
func (t TimezoneType) ToIANA() string {
	if iana, ok := timezoneIANAMap[t]; ok {
		return iana
	}
	return "UTC"
}

// Label 获取显示名称
func (t TimezoneType) Label() string {
	if label, ok := timezoneLabelMap[t]; ok {
		return label
	}
	return "未知时区"
}

// ToLocation 转换为time.Location
func (t TimezoneType) ToLocation() *time.Location {
	loc, err := time.LoadLocation(t.ToIANA())
	if err != nil {
		return time.UTC
	}
	return loc
}
