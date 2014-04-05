package core

//Core logic for Unlocking

import (
	"flag"
	"log"
	"time"
	"vnw/config"
	"vnw/gpio"
)

var uTime = flag.Int("utime", 20, "Number of seconds to unlock on successful swipe")

var Failed map[string]bool
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
	Failed = make(map[string]bool, 10)
}

func Auth(id string) {
	m := (*config.Cards)[id]
	log.Print("Saw card: ", id)
	log.Print("Mapped to member: ", m)
	if m != nil {
		m.Log(id)
		Unlock()
	} else {
		Failed[id] = true
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
