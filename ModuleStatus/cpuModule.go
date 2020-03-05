package ModuleStatus

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
	"time"
)

type CpuModule struct {
	AlarmStatus int
}

func (m *CpuModule) ModuleName() string {
	return "CPU使用率"
}

func (m *CpuModule) CheckInterval() time.Duration {
	return time.Millisecond * 6000
}

func (m *CpuModule) NormalStatus() string {
	return "65%"
}

func (m *CpuModule) CurrentStatus() (string, error) {
	data,err := cpu.Percent(time.Second * 5, false)
	return fmt.Sprintf("%d%%", int64(data[0])), err
}

func (m *CpuModule) IsExceedLimit(status string) bool {
	status0 := m.NormalStatus()
	v,_ := strconv.Atoi(status[:len(status)-1])
	v0,_ := strconv.Atoi(status0[:len(status0)-1])
	return v > v0
}

func (m *CpuModule) GetAlarmStatus() int {
	return m.AlarmStatus
}

func (m *CpuModule) SetAlarmStatus(status int) {
	m.AlarmStatus = status
}
