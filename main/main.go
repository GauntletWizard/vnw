package main

import (
	"flag"
	"log"
	"vnw/core"
	"vnw/config"
	"vnw/gpio"
	"vnw/nfc"
	"vnw/ui"
)

func init() {}

func main() {
	flag.StringVar(&ui.Httplistener, "port", ":80", "Listen Address for webserver")
	flag.StringVar(&gpio.Gpiodir, "gpiodir", "/sys/class/gpio", "Directory that holds GPIO pins. Exported for testing.")
	flag.IntVar(&gpio.Pin, "gpiopin", 60, "GPIO Pin to use")
	log.Print("Log message")
	flag.Parse()
	log.Print(gpio.Pin)
	gpio.Setup()
	
	log.Println("Starting config")
	config.Start()
	core.Start()
	nfc.Start()
	nfc.Poll()
	log.Println("Starting UI Server")
	ui.Start()
	log.Println("Shouldn't reach this")
}
