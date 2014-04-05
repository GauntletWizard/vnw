package core

//Core logic for Unlocking

import (
	"flag"
	"time"
	"vnw/config"
	"vnw/gpio"
  "log"
)

var uTime = flag.Int("utime", 20, "Number of seconds to unlock on successful swipe")

var failed []string
var doorTimer *time.Timer
var doorState bool
var fuTimer time.Time
var fLock bool
var flTimer time.Time
var fUnlock bool

func init() {
	Clear()
	doorState = false
	doorTimer = time.AfterFunc(time.Duration(0), Lock)
}

func Clear() {
	failed = make([]string, 10)
}

func Auth(id string) {
	m := (*config.Cards)[id]
	if m != nil {
		m.Log(id)
		Unlock()
	} else {
		failed = append(failed, id)
		failed = failed[1:]
	}
}

func ForceUnlock() {
	fUnlock = true
	Unlock()
}

func ForceLock() {
	fLock = true
	Lock()
}

func Unlock() {
	doorState = true
	Eval()
	doorTimer.Reset(time.Second * time.Duration(*uTime))
}

func Lock() {
	doorState = false
	Eval()
}

func Eval() {
	now := time.Now()
	if fLock {
		gpio.Lock()
		return
	}
	if fUnlock {
		gpio.Unlock()
		return
	}
	if flTimer.After(now) {
		gpio.Lock()
		return
	}
	if fuTimer.After(now) {
		gpio.Unlock()
		return
	}
	if doorState {
		gpio.Unlock()
		return
	}
	gpio.Lock()
	return
}
