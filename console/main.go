package main

import (
	"fmt"

	"github.com/roylee0704/gron"
	"github.com/roylee0704/gron/xtime"
)

type printJob struct{ Msg string }

func (p printJob) Run() {
	fmt.Println(p.Msg)
}

func main() {

	var (
		daily     = gron.Every(1 * xtime.Day)
		weekly    = gron.Every(1 * xtime.Week)
		monthly   = gron.Every(30 * xtime.Day)
		yearly    = gron.Every(365 * xtime.Day)
		purgeTask = func() { fmt.Println("purge unwanted records") }
		printFoo  = printJob{"Foo"}
		printBar  = printJob{"Bar"}
	)

	c := gron.New()

	c.AddFunc(weekly, func() { fmt.Println("Every week") })
	c.Add(daily.At("12:30"), printFoo)
	c.Start()

	// Jobs may also be added to a running Cron
	c.Add(monthly, printBar)
	c.AddFunc(yearly, purgeTask)

	// Stop the scheduler (does not stop any jobs already running).
	defer c.Stop()
}
