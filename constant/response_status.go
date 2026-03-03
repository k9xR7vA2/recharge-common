package constant

type GlobalResponseStatus string

const (
	ResponseStatusSuccess    NotifyStatus = "success"    //全局成功状态
	ResponseStatusFail       NotifyStatus = "fail"       //全局失败状态
	ResponseStatusProcessing NotifyStatus = "processing" //全局失败状态
)

func (b GlobalResponseStatus) String() string {
	return string(b)
}
