package main

import (
	"vnw/config"
	"vnw/gpio"
	//	"vnw/nfc"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"vnw/core"
	"vnw/ui"
)

var (
	secretFile = flag.String("secretfile", "", "Shared secret for grabbing member database.")
)

func main() {
	flag.StringVar(&ui.Httplistener, "port", ":8080", "Listen Address for webserver")
	u, _ := user.Current()
	flag.StringVar(&gpio.Gpiodir, "gpiodir", u.HomeDir+"/gpio", "Directory that holds GPIO pins. Exported for testing.")
	flag.StringVar(&config.File, "dbfile", "foo.csv", "location to read/store the user database")
	flag.StringVar(&config.Reqpath, "reqpath", "http://tcbtech.org/~ted/stuff/foo.csv", "URL of member list")
	flag.IntVar(&config.Sleep, "sleeptime", 600, "Number of seconds between updates of configfile")
	_ = flag.String("log", "", "Logfile. Inactive.")
	fmt.Println("Log message")
	flag.Parse()

	if *secretFile != "" {
		var err error
		log.Print("Loading secret from file")
		config.Secret, err = ioutil.ReadFile(*secretFile)
		if err != nil {
			log.Fatal("Unable to load secret file: ", err)
		}
	}

	gpio.Setup()
	fmt.Println("Starting config")
	config.Start()
	//	nfc.Poll()
	core.Start()
	fmt.Println("Starting UI Server")
	ui.Start()
	fmt.Println("Shouldn't reach this")
}
