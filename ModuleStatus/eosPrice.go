package ModuleStatus

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"net/http"
	"time"
)

type EOSPrice struct {
	AlarmStatus int
}

func (m *EOSPrice) ModuleName() string {
	return "CPU使用率"
}

func (m *EOSPrice) CheckInterval() time.Duration {
	return time.Hour
}

func (m *EOSPrice) NormalStatus() string {
	return "2.5"
}

func (m *EOSPrice) CurrentStatus() (string, error) {
	resp, err := http.Get("http://api.coincap.io/v2/assets?ids=eos")
	if err != nil {
		return "0",err
	}

	defer resp.Body.Close()
	js,_ := simplejson.NewFromReader(resp.Body)
	arr,_ := js.Get("data").Array()
	mp := arr[0].(map[string]interface{})

	return fmt.Sprintf("%s", mp["priceUsd"].(string)),nil
}

func (m *EOSPrice) IsExceedLimit(status string) bool {
	return status < m.NormalStatus()
}

func (m *EOSPrice) GetAlarmStatus() int {
	return 0
}

func (m *EOSPrice) SetAlarmStatus(status int) {
	m.AlarmStatus = status
}
