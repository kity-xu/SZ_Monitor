package main 

import (
	l4g "github.com/alecthomas/log4go"
	//"monitor.im/remote/protocol"
	"monitor.im/local/protocol"
)

func main() {
	// l4g.Info("in main StartPostServe top...")
	// protocol.NewHttpServer().StartPostServer()
	
	// l4g.Info("in main StartPostServe end...")

	l4g.Info("in main StartMonitor top...")
	pro := new (protocol.MonitorCode)
	pro.StartMonitor()
	
}