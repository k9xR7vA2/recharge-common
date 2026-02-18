package constant

type NotifyStatus string

const (
	NotifySuccess NotifyStatus = "SUCCESS" //全局成功状态
	NotifyFail    NotifyStatus = "FAIL"    //全局失败状态
)

func (b NotifyStatus) String() string {
	return string(b)
}
