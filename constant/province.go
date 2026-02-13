package constant

type ProvinceTag struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

var ProvinceList = []ProvinceTag{
	{Value: 200, Label: "广东"},
	{Value: 270, Label: "湖北"},
	{Value: 351, Label: "山西"},
	{Value: 471, Label: "内蒙古"},
	{Value: 951, Label: "宁夏"},
	{Value: 280, Label: "四川"},
	{Value: 290, Label: "陕西"},
	{Value: 431, Label: "吉林"},
	{Value: 551, Label: "安徽"},
	{Value: 571, Label: "浙江"},
	{Value: 898, Label: "海南"},
	{Value: 311, Label: "河北"},
	{Value: 371, Label: "河南"},
	{Value: 531, Label: "山东"},
	{Value: 871, Label: "云南"},
	{Value: 220, Label: "天津"},
	{Value: 731, Label: "湖南"},
	{Value: 931, Label: "甘肃"},
	{Value: 230, Label: "重庆"},
	{Value: 591, Label: "福建"},
	{Value: 771, Label: "广西"},
	{Value: 851, Label: "贵州"},
	{Value: 971, Label: "青海"},
	{Value: 100, Label: "北京"},
	{Value: 210, Label: "上海"},
	{Value: 240, Label: "辽宁"},
	{Value: 250, Label: "江苏"},
	{Value: 451, Label: "黑龙江"},
	{Value: 791, Label: "江西"},
	{Value: 891, Label: "西藏"},
	{Value: 991, Label: "新疆"},
}
