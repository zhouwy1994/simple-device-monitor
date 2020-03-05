package ModuleStatus

import "time"

// ModuleStatus 模块状态接口
type ModuleStatus interface {
	// 监控模块名
	ModuleName() string
	// 建议检测间隔
	CheckInterval() time.Duration
	// 模块状态最大正常值
	NormalStatus() string
	// 当前模块状态值
	CurrentStatus() (string, error)
	// 模块状态值是否超过正常范围
	IsExceedLimit(status string) bool
	// 告警状态(0:未发送任何信 1:已发送报警信息 2:已发送恢复信息)
	GetAlarmStatus() int
	// 设置告警状态
	SetAlarmStatus(status int)
}