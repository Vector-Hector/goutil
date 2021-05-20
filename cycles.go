package util

import (
	"sync"
	"time"
)

// Starts a cycle, that calls the toCall function after the given interval (seconds). If you want to keep it running, add the group and wait for the group to be done (will never happen)
func StartCycle(group *sync.WaitGroup, interval int, toCall func()) {
	if group != nil {
		group.Add(1)
	}
	go syncedCycle(group, interval, toCall)
}

func syncedCycle(group *sync.WaitGroup, interval int, toCall func()) {
	if group != nil {
		defer group.Done()
	}

	cycle(interval, toCall)
}

// TODO: add persistent error detection

func cycle(interval int, toCall func()) {
	defer cycle(interval, toCall)
	defer HandleUpdaterErrors()

	wait(interval)
	toCall()
}

func wait(seconds int) {
	time.Sleep(time.Duration(int(time.Second) * seconds))
}

