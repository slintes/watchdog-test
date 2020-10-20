package main

import (
	"github.com/coreos/go-systemd/daemon"
	"os"
	"strconv"
	"time"
)

func main() {


	healthy := true

	// get the service's watchdof timeout
	watchdogRefresh := time.Second
	if watchdogTimeout := os.Getenv("WATCHDOG_USEC"); watchdogTimeout != "" {
		println("WATCHDOG_USEC: ", watchdogTimeout)
		if t, err := strconv.Atoi(watchdogTimeout); err == nil {
			// seems to be good practice to use watchdog / 2
			// note that the watchdog env var's time unit is microseconds
			watchdogRefresh = time.Duration(t/2) * time.Microsecond
		} else {
			println(err)
		}
	} else {
		println("WATCHDOG_USEC not set?!")
	}
	println("watchog refesh: ", watchdogRefresh)

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

	for i:=0; i<3; i++ {
		println("healthy... ", i)
		time.Sleep(5 * time.Second)
	}
	// simulate an error
	healthy = false
	for i:=0; i<3; i++ {
		println("unhealthy... ", i)
		time.Sleep(5 * time.Second)
	}
	println("done without being restarted?!")

	daemon.SdNotify(false, daemon.SdNotifyStopping)

}
