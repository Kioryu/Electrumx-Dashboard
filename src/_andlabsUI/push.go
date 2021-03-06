package _andlabsUI

import (
	"github.com/sparrc/go-ping"
	"time"
	"github.com/andlabs/ui"
)

var pingBool bool = false

func serverPing(url string, laber *ui.Label) {
	pinger, err := ping.NewPinger(url)
	if checkError(err, false) {return}
	pinger.SetPrivileged(true)
	pinger.Timeout = time.Duration(time.Second * 1)
	pinger.Count = 1
	pinger.Run()
	stats := pinger.Statistics()
	if len(stats.Rtts) == 0 {
		boardLog.writeInfo("Server false")
		laber.SetText("Server false")
	} else {
		boardLog.writeInfo("Server true")
		laber.SetText("Server true")
	}
}
