package action

import (
	"os/exec"
	"sync"

	"github.com/go-ping/ping"
	"github.com/ritarock/groupping/lib/cui"
	groupping "github.com/ritarock/groupping/lib/ping"
)

func Run(targets []string) {
	var successTarget []string
	var errorTarget []string
	var wg sync.WaitGroup
	var result []*ping.Statistics

	for _, v := range targets {
		_, err := exec.Command("host", v).Output()
		if err != nil {
			errorTarget = append(errorTarget, v)
			continue
		}
		successTarget = append(successTarget, v)
	}

	for _, v := range successTarget {
		wg.Add(1)
		go func(v string) {
			result = append(result, groupping.DoPing(v))
			wg.Done()
		}(v)
	}
	wg.Wait()

	cui.View(result, errorTarget)
}
