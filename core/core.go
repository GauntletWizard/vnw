package core

//Core logic for Unlocking

import (
	"log"
	"time"
	"vnw/config"
	"vnw/gpio"
)

var UTime int
var Auth chan<- string
var Failed map[string]bool
var doorTimer *time.Timer
var doorState bool
var fuTimer time.Time
var fLock bool
var flTimer time.Time
var fUnlock bool

func Start() {
	Clear()
	doorState = false
	doorTimer = time.AfterFunc(time.Duration(0), Lock)
	a := make(chan string, 0)
	go auth(a)
	Auth = a
}

func Clear() {
	Failed = make(map[string]bool, 10)
}

func auth(a <-chan string) {
	for id := <-a; ; id = <-a {
		m := (*config.Cards)[id]
		log.Print("Saw card: ", id)
		log.Print("Mapped to member: ", m)
		if m != nil {
			go m.Log(id)
			Unlock()
		} else {
			Failed[id] = true
		}
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
	//	doorTimer.Stop()
	//	doorTimer = time.AfterFunc(time.Duration(time.Second * time.Duration(UTime)), Lock)
	doorTimer.Reset(time.Second * time.Duration(UTime))
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
