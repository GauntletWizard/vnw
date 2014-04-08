package main

import "vnw/gpio"
import "flag"

func main() {
	flag.StringVar(&gpio.Gpiodir, "gpiodir", "/sys/class/gpio", "Directory that holds GPIO pins. Exported for testing.")
	flag.IntVar(&gpio.Pin, "gpiopin", 60, "GPIO Pin to use")
	flag.Parse()
	gpio.High()
}
