// Drop dead simple GPIO binding for
package gpio

import (
	"log"
	"os"
	"strconv"
)

var Gpiodir string
var Pin int

func Setup() {
	f, err := os.OpenFile(Gpiodir + "/export", os.O_WRONLY, 0)
	if err != nil {
		log.Fatal("Could not open GPIO exports file!", err)
	}
	f.WriteString(strconv.Itoa(Pin))
}

func open() *os.File {
	f, err := os.Create(Gpiodir + "/gpio" + strconv.Itoa(Pin) + "/direction")
	if err != nil {
		log.Panic("Could not manipulate GPIO Pin "+strconv.Itoa(Pin), err)
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
	f, err := os.Open(Gpiodir + "/gpio" + strconv.Itoa(Pin) + "/value")
	if err != nil {
		log.Fatal("Could not manipulate GPIO Pin "+strconv.Itoa(Pin), err)
	}
	a := make([]byte, 1)
	f.Read(a)
	b, err := strconv.ParseBool(string(a))
	return b
}
