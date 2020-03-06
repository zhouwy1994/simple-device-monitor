package AlarmNotifier

import (
	"fmt"
	"github.com/go-gomail/gomail"
)

// EmailNotifier 邮箱报警通知器
type EmailNotifier struct {
	User string // 邮箱地址
	Pass string // 登陆密码(授权码)
	Host string // 邮箱服务器主机地址
	Port int // 邮箱服务器主机端口
	SendTo []string // 报警信息发送地址
}

func NewEmailNotifier(user, pass, host string, port int, sendTo []string) *EmailNotifier {
	return &EmailNotifier{
		User:user,
		Pass:pass,
		Host:host,
		Port:port,
		SendTo:sendTo,
	}
}
func (e *EmailNotifier) Notify(modName, modNormalStatus, modCurrentStatus, otherMessage , status string) error {
	m := gomail.NewMessage()
	m.SetHeader("From","Zhouwy" + "<" + e.User + ">")
	m.SetHeader("To", e.SendTo...)
	m.SetHeader("Subject", fmt.Sprintf("%s%s！！！", modName, status))
	m.SetBody("text/html", fmt.Sprintf(`%s%s:<br \>正常值:%s<br \>目前值:%s<br \>其他信息:%s<br \>`,
		modName, status, modNormalStatus, modCurrentStatus, otherMessage))
	return gomail.NewDialer(e.Host, e.Port, e.User, e.Pass).DialAndSend(m)
}

func (e *EmailNotifier)NotifyAlarm(modName, modNormalStatus, modCurrentStatus, otherMessage string) error {
	return e.Notify(modName, modNormalStatus, modCurrentStatus, otherMessage, "报警")
}

func (e *EmailNotifier)NotifyRecover(modName, modNormalStatus, modCurrentStatus, otherMessage string) error {
	return e.Notify(modName, modNormalStatus, modCurrentStatus, otherMessage, "恢复")
}