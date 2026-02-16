package traceConstant

const (
	// 基础HTTP信息
	KeyHTTPMethod    = "http.method"
	KeyHTTPURL       = "http.url"
	KeyHTTPPath      = "http.path" // URL路径部分
	KeyHTTPHost      = "http.host" // 主机名
	KeyClientIP      = "http.client_ip"
	KeyUserAgent     = "http.user_agent"     // 用户代理
	KeyContentType   = "http.content_type"   // 内容类型
	KeyContentLength = "http.content_length" // 内容长度
	KeyQueryString   = "http.query_string"   // URL查询参数
	KeyReferer       = "http.referer"        // 来源页面
	// 请求头部信息
	KeyRequestID     = "http.request_id"    // 请求ID(如果有)
	KeyAuthorization = "http.authorization" // 授权信息(注意敏感信息处理)
	// 性能相关
	KeyRequestTime = "http.request_time" // 请求时间
)

// 链路日志spanNode节点的Filed
const (
	// 错误相关
	KeyErrorNode    = "error.node"    // 错误发生的节点
	KeyErrorMessage = "error.message" // 错误信息
	KeyErrorParams  = "error.params"  // 错误相关参数
	KeyErrorStack   = "error.stack"   // 错误堆栈(如果有)
)
