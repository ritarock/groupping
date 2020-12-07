package cui

import (
	"fmt"
	"log"

	"github.com/go-ping/ping"
	"github.com/jroimartin/gocui"
)

type Widget struct {
	name           string
	body           string
	x0, y0, x1, y1 int
}

func (w *Widget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x0, w.y0, w.x1, w.y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, w.body)
	}
	return nil
}

func flowLayout(g *gocui.Gui) error {
	views := g.Views()
	x := 0
	for _, v := range views {
		w, h := v.Size()
		_, err := g.SetView(v.Name(), x, 0, x+w+1, h+1)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		x += w + 2
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func createComponent(pingResult []*ping.Statistics, errorTarget []string, maxX int, maxY int) []gocui.Manager {
	var widgets []gocui.Manager
	var successText string
	errorText := "Error Target\n"

	for _, v := range pingResult {
		successText = fmt.Sprintf("Target: %v\n 1stRtts: %v\n 2ndRtts: %v\n 3rdRtts: %v\n MaxRtt: %v\n MinRtt: %v\n AvgRtt: %v\n StdDevRtt: %v",
			v.Addr, v.Rtts[0], v.Rtts[1], v.Rtts[2], v.MaxRtt, v.MinRtt, v.AvgRtt, v.StdDevRtt)

		widgets = append(widgets, &Widget{v.Addr, successText, 0, 0, maxX/len(pingResult) - 1, maxY/2 + 3})
	}
	widgets = append(widgets, gocui.ManagerFunc(flowLayout))

	for _, v := range errorTarget {
		errorText += "Error: " + v + "\n"
	}
	widgets = append(widgets, &Widget{"Error", errorText, 0, maxY/2 + 4, maxX - 1, maxY - 1})

	return widgets
}

func View(pingResult []*ping.Statistics, errorTarget []string) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	maxX, maxY := g.Size()

	component := createComponent(pingResult, errorTarget, maxX, maxY)
	g.SetManager(component...)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
