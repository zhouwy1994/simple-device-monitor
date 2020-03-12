package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

import (
	af "github.com/zhouwy1994/simple-device-monitor/AlarmNotifier"
	ms "github.com/zhouwy1994/simple-device-monitor/ModuleStatus"
)

func newTickerFunc(d time.Duration, md ms.ModuleStatus,
	mds []ms.ModuleStatus, nr af.AlarmNotifier) func() {
	tk := time.NewTicker(d)
	ctx,cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-tk.C:
				checkStatus(md, mds, nr)
			case <-ctx.Done():
				return
			}
		}
	}()

	return func() {
		tk.Stop()
		cancel()
	}
}

func generateAllModStatusMessage(mds []ms.ModuleStatus) string {
	msg := ""
	for _, md := range mds {
		status,_ := md.CurrentStatus()
		msg += fmt.Sprintf(`<br \>%s:%s`, md.ModuleName(), status)
	}

	return msg
}

func checkStatus(md ms.ModuleStatus,
	mds []ms.ModuleStatus, nr af.AlarmNotifier) {
	modulestatus,_ := md.CurrentStatus()
	alarmStatus := md.GetAlarmStatus()
	if md.IsExceedLimit(modulestatus) && alarmStatus != 1 {
		md.SetAlarmStatus(1)
		otherMsg := generateAllModStatusMessage(mds)
		nr.NotifyAlarm(md.ModuleName(), md.NormalStatus(), modulestatus, otherMsg)
	} else if alarmStatus == 1 {
		md.SetAlarmStatus(2)
		otherMsg := generateAllModStatusMessage(mds)
		nr.NotifyRecover(md.ModuleName(), md.NormalStatus(), modulestatus, otherMsg)
	}
}

func main() {
	enr := af.NewEmailNotifier("xxx@aliyun.com", "xxx",
		"smtp.aliyun.com",25, []string{"xxx@foxmail.com"})

	mds := []ms.ModuleStatus{&ms.EOSPrice{}}

	stopFuncs := make([]func(), 0)
	for _,md := range mds {
		sf := newTickerFunc(md.CheckInterval(), md, mds, enr)
		stopFuncs = append(stopFuncs, sf)
	}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	select {
	case <-sigterm:
	}

	for _,f := range stopFuncs { f() }
}
