package constant

type AccountType int

// 1租金账户，2授信账户
const (
	RentalAccount AccountType = iota + 1 //1租金账户
	CreditAccount                        //2授信账户
)

func (b AccountType) Value() int {
	switch b {
	case RentalAccount:
		return 1
	case CreditAccount:
		return 2
	default:
		return 0
	}
}

func (b AccountType) Text() string {
	switch b {
	case RentalAccount:
		return "租金账户"
	case CreditAccount:
		return "授信账户"
	default:
		return "未知账户"
	}
}

func (b AccountType) Code() string {
	switch b {
	case RentalAccount:
		return "rental"
	case CreditAccount:
		return "credit"
	default:
		return "unknown"
	}
}

func (b AccountType) IsValid() bool {
	return b >= RentalAccount && b <= CreditAccount
}
