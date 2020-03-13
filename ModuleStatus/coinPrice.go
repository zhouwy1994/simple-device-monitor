package ModuleStatus

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"net/http"
	"time"
)

type CoinPrice struct {
	CoinName string
}

func (m *CoinPrice) ModuleName() string {
	return "自选币价格"
}

func (m *CoinPrice) CheckInterval() time.Duration {
	return time.Hour
}

func (m *CoinPrice) NormalStatus() string {
	return "2.5"
}

func (m *CoinPrice) CurrentStatus() (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.coincap.io/v2/assets?ids=%s", m.CoinName))
	if err != nil {
		return "0",err
	}

	defer resp.Body.Close()
	js,_ := simplejson.NewFromReader(resp.Body)
	arr,_ := js.Get("data").Array()
	if len(arr) < 1 {
		return "0",nil
	}

	mp := arr[0].(map[string]interface{})
	return fmt.Sprintf("%s", mp["priceUsd"].(string)),nil
}

func (m *CoinPrice) IsExceedLimit(status string) bool {
	return status < m.NormalStatus()
}

func (m *CoinPrice) GetAlarmStatus() int {
		return 0
}

func (m *CoinPrice) SetAlarmStatus(status int) {
}
