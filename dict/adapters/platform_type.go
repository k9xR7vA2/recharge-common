package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

// PlatformTypeDict 平台类型字典
type PlatformTypeDict struct{}

func (d *PlatformTypeDict) GetKey() string {
	return "platform_type"
}

func (d *PlatformTypeDict) GetName() string {
	return "平台类型"
}

func (d *PlatformTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		// 运营商
		{Label: constant.CuccPlatform.ShowName(), Value: string(constant.CuccPlatform), Code: string(constant.CuccPlatform)},
		{Label: constant.CmccPlatform.ShowName(), Value: string(constant.CmccPlatform), Code: string(constant.CmccPlatform)},
		{Label: constant.CtccPlatform.ShowName(), Value: string(constant.CtccPlatform), Code: string(constant.CtccPlatform)},
		// 电商
		{Label: constant.JdPlatform.ShowName(), Value: string(constant.JdPlatform), Code: string(constant.JdPlatform)},
		{Label: constant.TaoBaoPlatform.ShowName(), Value: string(constant.TaoBaoPlatform), Code: string(constant.TaoBaoPlatform)},
		{Label: constant.WoohooWebPlatform.ShowName(), Value: string(constant.WoohooWebPlatform), Code: string(constant.WoohooWebPlatform)},
		// 社交 App
		{Label: constant.DouyinAppPlatform.ShowName(), Value: string(constant.DouyinAppPlatform), Code: string(constant.DouyinAppPlatform)},
		{Label: constant.WeiboAppPlatform.ShowName(), Value: string(constant.WeiboAppPlatform), Code: string(constant.WeiboAppPlatform)},
		{Label: constant.XiaohongshuAppPlatform.ShowName(), Value: string(constant.XiaohongshuAppPlatform), Code: string(constant.XiaohongshuAppPlatform)},
		{Label: constant.KuaishouAppPlatform.ShowName(), Value: string(constant.KuaishouAppPlatform), Code: string(constant.KuaishouAppPlatform)},
		// 社交 Web
		{Label: constant.DouyinWebPlatform.ShowName(), Value: string(constant.DouyinWebPlatform), Code: string(constant.DouyinWebPlatform)},
		{Label: constant.WeiboWebPlatform.ShowName(), Value: string(constant.WeiboWebPlatform), Code: string(constant.WeiboWebPlatform)},
		{Label: constant.XiaohongshuWebPlatform.ShowName(), Value: string(constant.XiaohongshuWebPlatform), Code: string(constant.XiaohongshuWebPlatform)},
		{Label: constant.KuaishouWebPlatform.ShowName(), Value: string(constant.KuaishouWebPlatform), Code: string(constant.KuaishouWebPlatform)},
	}
}
