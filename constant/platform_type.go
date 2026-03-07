package constant

import "strings"

// PlatformType 社交平台类型
type PlatformType string

const (
	// 运营商
	CuccPlatform PlatformType = "cucc"
	CmccPlatform PlatformType = "cmcc"
	CtccPlatform PlatformType = "ctcc"

	// 电商
	JdPlatform     PlatformType = "jd"
	TaoBaoPlatform PlatformType = "taobao"

	// 社交 App
	DouyinAppPlatform      PlatformType = "douyin_app"
	WeiboAppPlatform       PlatformType = "weibo_app"
	XiaohongshuAppPlatform PlatformType = "xiaohongshu_app"
	KuaishouAppPlatform    PlatformType = "kuaishou_app"

	// 社交 Web
	DouyinWebPlatform      PlatformType = "douyin_web"
	WeiboWebPlatform       PlatformType = "weibo_web"
	XiaohongshuWebPlatform PlatformType = "xiaohongshu_web"
	KuaishouWebPlatform    PlatformType = "kuaishou_web"

	WoohooWebPlatform PlatformType = "woohoo_web"
)

func (p PlatformType) String() string {
	return string(p)
}

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
	case DouyinAppPlatform:
		return "抖音(App)"
	case WeiboAppPlatform:
		return "微博(App)"
	case XiaohongshuAppPlatform:
		return "小红书(App)"
	case KuaishouAppPlatform:
		return "快手(App)"
	case DouyinWebPlatform:
		return "抖音(Web)"
	case WeiboWebPlatform:
		return "微博(Web)"
	case XiaohongshuWebPlatform:
		return "小红书(Web)"
	case KuaishouWebPlatform:
		return "快手(Web)"
	case WoohooWebPlatform:
		return "WoohooWeb(Web)"
	default:
		return "未知平台"
	}
}

func (p PlatformType) IsValid() bool {
	switch p {
	case CuccPlatform, CmccPlatform, CtccPlatform,
		JdPlatform, TaoBaoPlatform,
		DouyinAppPlatform, WeiboAppPlatform, XiaohongshuAppPlatform, KuaishouAppPlatform,
		DouyinWebPlatform, WeiboWebPlatform, XiaohongshuWebPlatform, KuaishouWebPlatform, WoohooWebPlatform:
		return true
	default:
		return false
	}
}

func (p PlatformType) IsAppPlatform() bool {
	return strings.HasSuffix(string(p), "_app")
}

func (p PlatformType) IsWebPlatform() bool {
	return strings.HasSuffix(string(p), "_web")
}

// BasePlatform 获取基础平台名（去掉 _app/_web 后缀）
func (p PlatformType) BasePlatform() string {
	s := string(p)
	s = strings.TrimSuffix(s, "_app")
	s = strings.TrimSuffix(s, "_web")
	return s
}
