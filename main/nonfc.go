package main

import (
	"vnw/config"
	"vnw/gpio"
	//	"vnw/nfc"
	"flag"
	"fmt"
	"vnw/ui"
)

func init() {
	flag.StringVar(&ui.Httplistener, "port", ":80", "Listen Address for webserver")
	flag.StringVar(&gpio.Gpiodir, "gpiodir", "/sys/class/gpio", "Directory that holds GPIO pins. Exported for testing.")
	fmt.Println("Log message")
}

func main() {
	flag.Parse()
	gpio.Setup()
	fmt.Println("Starting config")
	config.Start()
	//	nfc.Poll()
	fmt.Println("Starting UI Server")
	ui.Start()
	fmt.Println("Shouldn't reach this")
}
