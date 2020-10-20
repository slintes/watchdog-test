package main

import (
	"time"

	"github.com/coreos/go-systemd/daemon"
)

func main() {

	healthy := true

	// get the service's watchdog timeout
	var timeout time.Duration
	var err error
	if timeout, err = daemon.SdWatchdogEnabled(false); err != nil {
		panic("watchdog not configured")
	}
	// recommendation is to notify at least at timeout / 2 period
	watchdogRefresh := timeout / 2
	println("watchdog refresh: ", watchdogRefresh)

	go func() {
		for {
			time.Sleep(watchdogRefresh)
			if healthy {
				println("notify")
				daemon.SdNotify(false, daemon.SdNotifyWatchdog)
			} else {
				println("no notify")
			}
		}
	}()

	daemon.SdNotify(false, daemon.SdNotifyReady)

	for i := 0; i < 3; i++ {
		println("healthy... ", i)
		time.Sleep(5 * time.Second)
	}
	// simulate an error
	healthy = false
	for i := 0; i < 3; i++ {
		println("unhealthy... ", i)
		time.Sleep(5 * time.Second)
	}
	println("done without being restarted?!")

	daemon.SdNotify(false, daemon.SdNotifyStopping)

}
