package main

import (
	"flag"
	"log"
	"os"
	"syscall"
	"os/signal"
	"vnw/config"
	"vnw/core"
	"vnw/gpio"
	"vnw/nfc"
	"vnw/ui"
)

func init() {}

var logf string

func main() {
	flag.StringVar(&ui.Httplistener, "port", ":80", "Listen Address for webserver")
	flag.StringVar(&gpio.Gpiodir, "gpiodir", "/sys/class/gpio", "Directory that holds GPIO pins. Exported for testing.")
	flag.IntVar(&gpio.Pin, "gpiopin", 7, "GPIO Pin to use")
	flag.IntVar(&core.UTime, "utime", 10, "Number of seconds to unlock on successful swipe")
	flag.IntVar(&config.Sleep, "sleeptime", 600, "Number of seconds between updates of configfile")
	flag.StringVar(&config.File, "dbfile", "foo.csv", "location to read/store the user database")
	flag.StringVar(&config.Reqpath, "reqpath", "http://tcbtech.org/~ted/stuff/foo.csv", "URL of member list")
	flag.StringVar(&config.Secret, "secret", "", "Shared secret for grabbing member database.")
	flag.StringVar(&logf, "log", "-", "File to output logging to. Defaults to Stdout")
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
	c := make(chan os.Signal, 0)
	go handleHup(c)
	signal.Notify(c, os.Signal(syscall.SIGHUP))
	ui.Start()
	log.Println("Shouldn't reach this")
}

func handleHup(c <-chan os.Signal) {
	for {
		var file *os.File
		if logf == "-" {
			file = os.Stdout
		} else {
			file, _ = os.Create(logf)
		}
		log.SetOutput(file)
		<-c
	}
}
