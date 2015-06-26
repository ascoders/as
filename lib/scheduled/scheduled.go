/*==================================================
	计划任务

	scheduled.Run(22, 32, 30, func() {
		fmt.Println("oh!")
	})

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package scheduled

import (
	"time"
)

type Scheduled struct{}

var (
	ScheduledInstance *Scheduled
)

func init() {
	ScheduledInstance = &Scheduled{}
}

func (this *Scheduled) Run(hour int, minute int, second int, callback func()) {
	go runningRoutine(hour, minute, second, callback)
}

func runningRoutine(hour int, minute int, second int, callback func()) {
	ticker := updateTicker(hour, minute, second)
	for {
		<-ticker.C
		// do something..
		callback()
		ticker = updateTicker(hour, minute, second)
	}
}

func updateTicker(hour int, minute int, second int) *time.Ticker {
	nextTick := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, time.Local)
	if !nextTick.After(time.Now()) {
		nextTick = nextTick.Add(24 * time.Hour)
	}
	diff := nextTick.Sub(time.Now())
	return time.NewTicker(diff)
}
