module github.com/zhouwy1994/simple-device-monitor

go 1.13

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/shirou/gopsutil v2.20.2+incompatible
	golang.org/x/sys v0.0.0-00010101000000-000000000000 // indirect
)

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20200302150141-5c8b2ff67527
