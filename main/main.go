package main

import (
	"vnw/config"
	"vnw/nfc"
	"vnw/ui"
)

func main() {
	config.Start()
	nfc.Poll()
	ui.Start()
}
