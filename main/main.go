package main

import (
	"flag"
	"log"
	"vnw/config"
	"vnw/core"
	"vnw/gpio"
	"vnw/nfc"
	"vnw/ui"
)

func init() {}

func main() {
	flag.StringVar(&ui.Httplistener, "port", ":80", "Listen Address for webserver")
	flag.StringVar(&gpio.Gpiodir, "gpiodir", "/sys/class/gpio", "Directory that holds GPIO pins. Exported for testing.")
	flag.IntVar(&gpio.Pin, "gpiopin", 60, "GPIO Pin to use")
	flag.IntVar(&core.UTime, "utime", 10, "Number of seconds to unlock on successful swipe")
	flag.IntVar(&config.Sleep, "sleeptime", 600, "Number of seconds between updates of configfile")
  flag.StringVar(&config.File, "dbfile", "foo.csv", "location to read/store the user database")
  flag.StringVar(&config.Reqpath, "reqpath", "http://tcbtech.org/~ted/stuff/foo.csv", "URL of member list")
	flag.StringVar(&config.Secret, "secret", "", "Shared secret for grabbing member database.")
	log.Print("Log message")
	flag.Parse()
	log.Print(gpio.Pin)
	gpio.Setup()

	log.Println("Starting config")
	config.Start()
	core.Start()
	nfc.Start()
	go nfc.Poll()
	log.Println("Starting UI Server")
	ui.Start()
	log.Println("Shouldn't reach this")
}
