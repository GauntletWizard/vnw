package nfc

import (
//"vnw/config"
"github.com/fuzxxl/nfc/latest/nfc"
"flag"
"log"
// "time"
)

var nfcdevice = flag.String("nfcdevice", "pn532_uart:/dev/ttyUSB0:115200", "LibNFC config string for NFC device to open")
var nfcpoll = flag.Int("nfcpoll", 100, "Miliseconds between NFC Polling routines")
var nfcwait = flag.Int("nfcwait", 2000, "Miliseconds to wait between a NFC event and next poll")

func Poll() {
  d, err := nfc.Open(*nfcdevice)
  if err != nil {
    log.Fatal("Failed to open NFC Device ", *nfcdevice, err)
  }
  nm := nfc.Modulation{
    Type: nfc.ISO14443A,
    BaudRate: nfc.NBR_106,
  }
  
  targets, err := d.InitiatorListPassiveTargets(nm)
  log.Print(targets)
}
