package constant

// PlatformType 社交平台类型
type PlatformType string

const (
	CuccPlatform        PlatformType = "cucc"
	CmccPlatform        PlatformType = "cmcc"
	CtccPlatform        PlatformType = "ctcc"
	JdPlatform          PlatformType = "jd"
	TaoBaoPlatform      PlatformType = "taobao"
	DouyinPlatform      PlatformType = "douyin"      // DouyinPlatform 抖音
	WeiboPlatform       PlatformType = "weibo"       // WeiboPlatform 微博
	XiaohongshuPlatform PlatformType = "xiaohongshu" // XiaohongshuPlatform 小红书
	KuaishouPlatform    PlatformType = "kuaishou"    // KuaishouPlatform 快手
)

func (p PlatformType) String() string {
	switch p {
	case CuccPlatform:
		return "cucc"
	case CmccPlatform:
		return "cmcc"
	case CtccPlatform:
		return "ctcc"
	case JdPlatform:
		return "jd"
	case TaoBaoPlatform:
		return "taobao"
	case DouyinPlatform:
		return "douyin"
	case WeiboPlatform:
		return "weibo"
	case XiaohongshuPlatform:
		return "xiaohongshu"
	case KuaishouPlatform:
		return "kuaishou"
	default:
		return "未知平台"
	}
}

// ShowName 实现匹配规则的字符串表示
func (p PlatformType) ShowName() string {
	switch p {
	case CuccPlatform:
		return "联通"
	case CmccPlatform:
		return "移动"
	case CtccPlatform:
		return "电信"
	case JdPlatform:
		return "京东"
	case TaoBaoPlatform:
		return "淘宝"
	case DouyinPlatform:
		return "抖音"
	case WeiboPlatform:
		return "微博"
	case XiaohongshuPlatform:
		return "小红书"
	case KuaishouPlatform:
		return "快手"
	default:
		return "未知平台"
	}
}

// IsValid 验证平台类型是否有效
func (p PlatformType) IsValid() bool {
	switch p {
	case CuccPlatform, CmccPlatform, CtccPlatform, JdPlatform, TaoBaoPlatform, DouyinPlatform, WeiboPlatform, XiaohongshuPlatform, KuaishouPlatform:
		return true
	default:
		return false
	}
}

// GetAllPlatformTypes 获取所有社交平台类型
func GetAllPlatformTypes() []struct {
	Label string       `json:"label"`
	Value PlatformType `json:"value"`
} {
	return []struct {
		Label string       `json:"label"`
		Value PlatformType `json:"value"`
	}{
		{Label: CuccPlatform.ShowName(), Value: CuccPlatform},
		{Label: CmccPlatform.ShowName(), Value: CmccPlatform},
		{Label: CtccPlatform.ShowName(), Value: CtccPlatform},
		{Label: JdPlatform.ShowName(), Value: JdPlatform},
		{Label: TaoBaoPlatform.ShowName(), Value: TaoBaoPlatform},

		{Label: DouyinPlatform.ShowName(), Value: DouyinPlatform},
		{Label: WeiboPlatform.ShowName(), Value: WeiboPlatform},
		{Label: XiaohongshuPlatform.ShowName(), Value: XiaohongshuPlatform},
		{Label: KuaishouPlatform.ShowName(), Value: KuaishouPlatform},
	}
}
