package constant

type GlobalResponseStatus string

const (
	ResponseStatusSuccess    GlobalResponseStatus = "success"    //全局成功状态
	ResponseStatusFail       GlobalResponseStatus = "fail"       //全局失败状态
	ResponseStatusProcessing GlobalResponseStatus = "processing" //全局失败状态
)

func (b GlobalResponseStatus) String() string {
	return string(b)
}
