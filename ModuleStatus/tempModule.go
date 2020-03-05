package ModuleStatus

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type TempModule struct {
	AlarmStatus int
}

func (m *TempModule) ModuleName() string {
	return "CPU温度"
}

func (m *TempModule) CheckInterval() time.Duration {
	return time.Millisecond * 5000
}

func (m *TempModule) NormalStatus() string {
	return "55°"
}

func (m *TempModule) CurrentStatus() (string, error) {
	data,err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return "",err
	}

	temp,err := strconv.Atoi(string(data[:len(data)-1]))
	if err != nil {
		return "",err
	}

	return fmt.Sprintf("%d°", temp / 1000),nil
}

func (m *TempModule) IsExceedLimit(status string) bool {
	status0 := m.NormalStatus()
	v,_ := strconv.Atoi(status[:len(status)-2])
	v0,_ := strconv.Atoi(status0[:len(status0)-2])
	return v > v0
}

func (m *TempModule) GetAlarmStatus() int {
	return m.AlarmStatus
}

func (m *TempModule) SetAlarmStatus(status int) {
	m.AlarmStatus = status
}
