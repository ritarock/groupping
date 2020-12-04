package ping

import (
	"github.com/go-ping/ping"
)

func DoPing(target string) *ping.Statistics {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
	return pinger.Statistics()
}
