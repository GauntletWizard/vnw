// Drop dead simple GPIO binding for 
package gpio

import (
  "flag"
  "log"
  "os"
  "strconv"
)

var fpin = flag.Int("gpiopin", 60, "GPIO Pin to use")
var pin int

func init() {
  flag.Parse()
  pin = *fpin
  f, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, 0)
  if err != nil {
    log.Fatal("Could not open GPIO exports file!", err)
  }
  f.WriteString(strconv.Itoa(pin))
}

func open() *os.File {
  f, err := os.Create("/sys/class/gpio/gpio" + strconv.Itoa(pin) + "/direction")
  if err != nil {
    log.Fatal("Could not manipulate GPIO pin " + strconv.Itoa(pin), err)
  }
  return f
}

func High() {
  f := open()
  f.WriteString("high")
}

func Low() {
  f := open()
  f.WriteString("low")
}

func Value() bool {
  f, err := os.Open("/sys/class/gpio/gpio" + strconv.Itoa(pin) + "/value")
  if err != nil {
    log.Fatal("Could not manipulate GPIO pin " + strconv.Itoa(pin), err)
  }
  a := make([]byte, 1)
  f.Read(a)
  b, err := strconv.ParseBool(string(a))
  return b
}
