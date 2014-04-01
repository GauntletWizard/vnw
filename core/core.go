package core

//Core logic for Unlocking

import (
"vnw/gpio"
"vnw/config"
)

var failed []string

func init() {
	failed = make([]string, 10)
}

func Auth(id string) {
	if config.
}

func ForceUnlock() {
	fUnlock = true
	Unlock()
}

func ForceLock()
	fLock = true
	Lock()
}

func Eval() {
	if (fUnlock && !fLock) && 
