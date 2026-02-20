package payloads

type APIOperation string

const (
	APIOperationCreate       APIOperation = "CREATE"
	APIOperationUpdate       APIOperation = "UPDATE"
	APIOperationDelete       APIOperation = "DELETE"
	APIOperationRefreshCache APIOperation = "REFRESH_CACHE"
	APIOperationSync         APIOperation = "SYNC"
)

type TenantSystemPermissionUpdateTask struct {
	Operation  APIOperation `json:"operation"`
	OperatorID uint         `json:"operator_id,omitempty"`
	Remark     string       `json:"remark,omitempty"`
	Timestamp  int64        `json:"timestamp"`
}

func GetOperationName(op APIOperation) string {
	switch op {
	case APIOperationCreate:
		return "创建API"
	case APIOperationUpdate:
		return "更新API"
	case APIOperationDelete:
		return "删除API"
	case APIOperationRefreshCache:
		return "刷新权限缓存"
	case APIOperationSync:
		return "同步API"
	default:
		return "未知操作"
	}
}
