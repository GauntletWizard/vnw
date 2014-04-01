package main

import (
	"vnw/config"
	"vnw/nfc"
	"vnw/ui"
)

func main() {
	config.Start()
	go nfc.Poll()
	ui.Start()
}
