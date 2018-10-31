package models

import "time"

const (
	TurnOnNode         = "TurnOnNode"         // 开机
	TurnOffNode        = "TurnOffNode"        // 关机
	RebootNode         = "RebootNode"         // 重启
	EnableNode         = "EnableNode"         // 启用
	DisableNode        = "DisableNode"        // 禁用
	ReleaseNode        = "ReleaseNode"        // 释放
	TrashNode          = "TrashNode"          // 删除
	RenameNode         = "RenameNode"         // 修改名称
	GetConsoleNode     = "GetConsoleNode"     // 获取控制台
	ChangeSystemNode   = "ChangeSystemNode"   // 修改系统
	ChangePasswordNode = "ChangePasswordNode" // 修改密码
	ActivateNode       = "ActivateNode"       // 开通
)

// 资源列表（json字符串）
type Resources struct {
	Id   string `json:"id"`   // 项目ID
	Name string `json:"name"` // 资源名称
	Type string `json:"type"` // 资源类型
}

// 资源日志 - 开始
type LogSourceStart struct {
	CreateTime   time.Time `json:"createTime"`   // 操作时间（精确到毫秒） yyyy-MM-DD HH:mm:ss.S
	AuditType    string    `json:"auditType"`    // 业务类型：resource（资源类）
	LogType      string    `json:"logType"`      // 日志类型（LOG_START）
	Action       string    `json:"action"`       // 操作action
	RequestId    string    `json:"requestId"`    // 操作请求ID
	Version      string    `json:"version"`      // 版本号
	VisitorId    string    `json:"visitorId"`    // 操作者ID
	CloudId      string    `json:"cloudId"`      // 云环境ID
	DataCenterId string    `json:"dataCenterId"` // 数据中心ID
	OwnerId      string    `json:"ownerId"`      // 项目ID
	OwnerName    string    `json:"ownerName"`    // 项目名称
}

// 资源日志 - 结束
type LogSourceEnd struct {
	CreateTime   time.Time   `json:"createTime"`   // 操作时间（精确到毫秒） yyyy-MM-DD HH:mm:ss.S
	AuditType    string      `json:"auditType"`    // 业务类型：resource（资源类）
	LogType      string      `json:"logType"`      // 日志类型（LOG_END）
	Action       string      `json:"action"`       // 操作action
	RequestId    string      `json:"requestId"`    // 操作请求ID
	Version      string      `json:"version"`      // 版本号
	VisitorId    string      `json:"visitorId"`    // 操作者ID
	CloudId      string      `json:"cloudId"`      // 云环境ID
	DataCenterId string      `json:"dataCenterId"` // 数据中心ID
	OwnerId      string      `json:"ownerId"`      // 项目ID
	OwnerName    string      `json:"ownerName"`    // 项目名称
	Resources    []Resources `json:"resources"`    // 资源列表（json字符串）
	Result       string      `json:"result"`       // 操作结果（failure、success）
}

// 业务日志 - 开始
type LogServiceStart struct {
	CreateTime time.Time `json:"createTime"` // 操作时间（精确到毫秒） yyyy-MM-DD HH:mm:ss.S
	AuditType  string    `json:"auditType"`  // business（业务类）
	LogType    string    `json:"logType"`    // 日志类型（LOG_START）
	Action     string    `json:"action"`     // 操作action
	RequestId  string    `json:"requestId"`  // 操作请求ID
	Version    string    `json:"version"`    // 版本号
	VisitorId  string    `json:"visitorId"`  // 操作者ID
	//CloudId      string    `json:"cloudId"`      // 云环境ID
	//DataCenterId string    `json:"dataCenterId"` // 数据中心ID
	OwnerId   string `json:"ownerId"`   // 项目ID
	OwnerName string `json:"ownerName"` // 项目名称
}

// 业务日志 - 结束
type LogServiceEnd struct {
	CreateTime time.Time `json:"createTime"` // 操作时间（精确到毫秒） yyyy-MM-DD HH:mm:ss.S
	AuditType  string    `json:"auditType"`  // 业务类型：business（业务类）
	LogType    string    `json:"logType"`    // 日志类型（LOG_END）
	Action     string    `json:"action"`     // 操作action
	RequestId  string    `json:"requestId"`  // 操作请求ID
	Version    string    `json:"version"`    // 版本号
	VisitorId  string    `json:"visitorId"`  // 操作者ID
	//CloudId      string    `json:"cloudId"`      // 云环境ID
	//DataCenterId string    `json:"dataCenterId"` // 数据中心ID
	OwnerId   string      `json:"ownerId"`   // 项目ID
	OwnerName string      `json:"ownerName"` // 项目名称
	Resources []Resources `json:"resources"` // 资源列表（json字符串）
	Result    string      `json:"result"`    // 操作结果（failure、success）
}

// 日志 - 详情（资源、业务）
type LogDetail struct {
	CreateTime   time.Time `json:"createTime"`   // 操作时间（精确到毫秒） yyyy-MM-DD HH:mm:ss.S
	LogType      string    `json:"logType"`      // 日志类型（LOG_DETAIL）
	Action       string    `json:"action"`       // 每一步详情的操作动作（由各服务自定义，语义清晰即可， 如create user、get project information等）
	ServiceName  string    `json:"serviceName"`  // 日志来源服务（注册中心的服务名），谁记录的日志
	RequestId    string    `json:"requestId"`    // 操作请求ID
	Version      string    `json:"version"`      // 版本号
	Params       string    `json:"params"`       // 当前详情操作涉及的参数（json字符串，由各服务自定义，数据要方便定位错误）
	ErrorCode    string    `json:"errorCode"`    // 错误码
	ErrorMessage string    `json:"errorMessage"` // 错误信息
	Result       string    `json:"result"`       // 操作结果（failure、success）
}

// 日志 - 登录
type LogLogin struct {
	RequestId      string    `json:"requestId"`      // 操作请求ID
	LoginTime      time.Time `json:"loginTime"`      // 操作时间（精确到毫秒） yyyy-MM-DD HH:mm:ss.S
	UserId         string    `json:"userId"`         // 用户Id
	UserName       string    `json:"userName"`       // 用户名
	UserType       string    `json:"userType"`       // 用户类型（local、ldap）
	LoginIp        string    `json:"loginIp"`        // 登录IP
	Result         string    `json:"result"`         // 登录结果（failure、success）
	FailureMessage string    `json:"failureMessage"` // 失败原因（result为failure时必传）
}
