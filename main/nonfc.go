package main

import (
	"vnw/config"
	"vnw/gpio"
	//	"vnw/nfc"
	"flag"
	"fmt"
	"os/user"
	"vnw/core"
	"vnw/ui"
)

func init() {
	flag.StringVar(&ui.Httplistener, "port", ":8080", "Listen Address for webserver")
	u, _ := user.Current()
	flag.StringVar(&gpio.Gpiodir, "gpiodir", u.HomeDir+"/gpio", "Directory that holds GPIO pins. Exported for testing.")
	flag.IntVar(&config.Sleep, "sleeptime", 600, "Number of seconds between updates of configfile")
	flag.StringVar(&config.Secret, "secret", "", "Shared secret for grabbing member database.")
	fmt.Println("Log message")
}

func main() {
	flag.Parse()
	gpio.Setup()
	fmt.Println("Starting config")
	config.Start()
	//	nfc.Poll()
	core.Start()
	fmt.Println("Starting UI Server")
	ui.Start()
	fmt.Println("Shouldn't reach this")
}
