// Drop dead simple GPIO binding for
package gpio

import (
	"flag"
	"log"
	"os"
	"strconv"
)

var fpin = flag.Int("gpiopin", 60, "GPIO Pin to use")
var Gpiodir string
var pin int

func Setup() {
	pin = *fpin
	f, err := os.OpenFile(Gpiodir + "/export", os.O_WRONLY, 0)
	if err != nil {
		log.Fatal("Could not open GPIO exports file!", err)
	}
	f.WriteString(strconv.Itoa(pin))
}

func open() *os.File {
	f, err := os.Create(Gpiodir + "/gpio" + strconv.Itoa(pin) + "/direction")
	if err != nil {
		log.Fatal("Could not manipulate GPIO pin "+strconv.Itoa(pin), err)
	}
	return f
}

func High() {
	f := open()
	f.WriteString("high")
}

func Unlock() {
	High()
}

func Low() {
	f := open()
	f.WriteString("low")
}

func Lock() {
	Low()
}

func Value() bool {
	f, err := os.Open(Gpiodir + "/gpio" + strconv.Itoa(pin) + "/value")
	if err != nil {
		log.Fatal("Could not manipulate GPIO pin "+strconv.Itoa(pin), err)
	}
	a := make([]byte, 1)
	f.Read(a)
	b, err := strconv.ParseBool(string(a))
	return b
}
