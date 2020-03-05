package AlarmNotifier

type AlarmNotifier interface {
	// 告警信息发送
	NotifyAlarm(modName, modNormalStatus, modCurrentStatus, otherMessage string) error
	// 恢复信息发送
	NotifyRecover(modName, modNormalStatus, modCurrentStatus, otherMessage string) error
}