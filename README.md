# A very simple Go based systemd watchdog implementation

## Usage

    go install
    sudo systemctl link ./watchdog-test.service
    sudo systemctl start watchdog-test.service
    journalctl -f -u watchdog-test.service
    
## What to expect

The app should be restarted by systemd over and over again:

```
Oct 20 17:32:40 slintes systemd[1]: Starting Watchdog Test service...
Oct 20 17:32:40 slintes watchdog-test[2164070]: WATCHDOG_USEC:  5000000
Oct 20 17:32:40 slintes watchdog-test[2164070]: watchog refesh:  2500000000
Oct 20 17:32:40 slintes systemd[1]: Started Watchdog Test service.
Oct 20 17:32:40 slintes watchdog-test[2164070]: healthy...  0
Oct 20 17:32:42 slintes watchdog-test[2164070]: notify
Oct 20 17:32:45 slintes watchdog-test[2164070]: notify
Oct 20 17:32:45 slintes watchdog-test[2164070]: healthy...  1
Oct 20 17:32:47 slintes watchdog-test[2164070]: notify
Oct 20 17:32:50 slintes watchdog-test[2164070]: notify
Oct 20 17:32:50 slintes watchdog-test[2164070]: healthy...  2
Oct 20 17:32:52 slintes watchdog-test[2164070]: notify
Oct 20 17:32:55 slintes watchdog-test[2164070]: unhealthy...  0
Oct 20 17:32:55 slintes watchdog-test[2164070]: no notify
Oct 20 17:32:57 slintes watchdog-test[2164070]: no notify
Oct 20 17:32:58 slintes systemd[1]: watchdog-test.service: Watchdog timeout (limit 5s)!
Oct 20 17:32:58 slintes systemd[1]: watchdog-test.service: Killing process 2164070 (watchdog-test) with signal SIGABRT.
Oct 20 17:32:58 slintes watchdog-test[2164070]: SIGABRT: abort
Oct 20 17:32:58 slintes watchdog-test[2164070]: PC=0x469001 m=0 sigcode=0
Oct 20 17:32:58 slintes watchdog-test[2164070]: goroutine 0 [idle]:
Oct 20 17:32:58 slintes watchdog-test[2164070]: runtime.futex(0x585bc8, 0x80, 0x0, 0x0, 0x7fff00000000, 0x468bdc, 0x67bcc, 0x513af0e, 0x7fff1437b338, 0x40bb3f, ...)
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/sys_linux_amd64.s:587 +0x21
Oct 20 17:32:58 slintes watchdog-test[2164070]: runtime.futexsleep(0x585bc8, 0x0, 0xffffffffffffffff)
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/os_linux.go:45 +0x46
Oct 20 17:32:58 slintes watchdog-test[2164070]: runtime.notesleep(0x585bc8)
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/lock_futex.go:159 +0x9f
Oct 20 17:32:58 slintes watchdog-test[2164070]: runtime.stopm()
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/proc.go:1910 +0xc5
Oct 20 17:32:58 slintes watchdog-test[2164070]: runtime.findrunnable(0xc00001e000, 0x0)
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/proc.go:2471 +0xa7f
Oct 20 17:32:58 slintes watchdog-test[2164070]: runtime.schedule()
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/proc.go:2669 +0x2d7
Oct 20 17:32:58 slintes watchdog-test[2164070]: runtime.park_m(0xc000001980)
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/proc.go:2837 +0x9d
Oct 20 17:32:58 slintes watchdog-test[2164070]: runtime.mcall(0x0)
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/asm_amd64.s:318 +0x5b
Oct 20 17:32:58 slintes watchdog-test[2164070]: goroutine 1 [sleep]:
Oct 20 17:32:58 slintes watchdog-test[2164070]: time.Sleep(0x12a05f200)
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/time.go:188 +0xbf
Oct 20 17:32:58 slintes watchdog-test[2164070]: main.main()
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/dev/work/lifecycle/watchdog-test/main.go:50 +0x249
Oct 20 17:32:58 slintes watchdog-test[2164070]: goroutine 6 [sleep]:
Oct 20 17:32:58 slintes watchdog-test[2164070]: time.Sleep(0x9502f900)
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/.gimme/versions/go1.15.linux.amd64/src/runtime/time.go:188 +0xbf
Oct 20 17:32:58 slintes watchdog-test[2164070]: main.main.func1(0x9502f900, 0xc000016234)
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/dev/work/lifecycle/watchdog-test/main.go:30 +0x2f
Oct 20 17:32:58 slintes watchdog-test[2164070]: created by main.main
Oct 20 17:32:58 slintes watchdog-test[2164070]:         /home/msluiter/dev/work/lifecycle/watchdog-test/main.go:28 +0x169
Oct 20 17:32:58 slintes watchdog-test[2164070]: rax    0xca
Oct 20 17:32:58 slintes watchdog-test[2164070]: rbx    0x585a80
Oct 20 17:32:58 slintes watchdog-test[2164070]: rcx    0x469003
Oct 20 17:32:58 slintes watchdog-test[2164070]: rdx    0x0
Oct 20 17:32:58 slintes watchdog-test[2164070]: rdi    0x585bc8
Oct 20 17:32:58 slintes watchdog-test[2164070]: rsi    0x80
Oct 20 17:32:58 slintes watchdog-test[2164070]: rbp    0x7fff1437b300
Oct 20 17:32:58 slintes watchdog-test[2164070]: rsp    0x7fff1437b2b8
Oct 20 17:32:58 slintes watchdog-test[2164070]: r8     0x0
Oct 20 17:32:58 slintes watchdog-test[2164070]: r9     0x0
Oct 20 17:32:58 slintes watchdog-test[2164070]: r10    0x0
Oct 20 17:32:58 slintes watchdog-test[2164070]: r11    0x286
Oct 20 17:32:58 slintes watchdog-test[2164070]: r12    0x3
Oct 20 17:32:58 slintes watchdog-test[2164070]: r13    0x585520
Oct 20 17:32:58 slintes watchdog-test[2164070]: r14    0x28
Oct 20 17:32:58 slintes watchdog-test[2164070]: r15    0x200
Oct 20 17:32:58 slintes watchdog-test[2164070]: rip    0x469001
Oct 20 17:32:58 slintes watchdog-test[2164070]: rflags 0x286
Oct 20 17:32:58 slintes watchdog-test[2164070]: cs     0x33
Oct 20 17:32:58 slintes watchdog-test[2164070]: fs     0x0
Oct 20 17:32:58 slintes watchdog-test[2164070]: gs     0x0
Oct 20 17:32:58 slintes systemd[1]: watchdog-test.service: Main process exited, code=exited, status=2/INVALIDARGUMENT
Oct 20 17:32:58 slintes systemd[1]: watchdog-test.service: Failed with result 'watchdog'.
Oct 20 17:32:59 slintes systemd[1]: watchdog-test.service: Scheduled restart job, restart counter is at 8.
Oct 20 17:32:59 slintes systemd[1]: Stopped Watchdog Test service.
Oct 20 17:32:59 slintes systemd[1]: Starting Watchdog Test service...
...
```

## Don't forget...

.. to stop it :)

    sudo systemctl stop watchdog-test.service
