package entities

type MobilePoolArgs struct {
	Amount      string // 订单金额
	Carrier     string // 运营商
	ChargeSpeed string // 充值速度
	Region      string // 区域
	Province    string // 省份，
}

// AddMobilePoolArgs 添加话费订单入池池参数
type AddMobilePoolArgs struct {
	SupplierOrderStr string
	TenantID         uint
	SystemOrderSn    string
	ValidTime        int
	ExpiredAt        int64
	//SupplierID       uint
	SupplierOrderSn string
	//CreatedAt        int64
	Priority string //优先级别 (high/normal)

	MobilePoolArgs
}

// 话费撮合参数
type MobileMatchmakingArgs struct {
	TenantID uint
	MobilePoolArgs
}
