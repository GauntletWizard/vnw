package nfc

import (
//"vnw/config"
"github.com/fuzxxl/nfc/latest/nfc"
"flag"
"log"
// "time"
"unsafe"
"encoding/hex"
)

/* 
#cgo LDFLAGS: -lnfc
#include <stdlib.h>
#include "nfc.h"
*/
import "C"

var nfcdevice = flag.String("nfcdevice", "pn532_uart:/dev/ttyUSB0:115200", "LibNFC config string for NFC device to open")
var nfcpoll = flag.Int("nfcpoll", 100, "Miliseconds between NFC Polling routines")
var nfcwait = flag.Int("nfcwait", 2000, "Miliseconds to wait between a NFC event and next poll")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	c := C.int(0)
	cs := C.CString(*nfcdevice)
	C.init_globals(cs)
	C.free(unsafe.Pointer(cs))
	t := C.get_ids(0xFF, c);
	log.Print(c, t);
	t1 := C.GoString(*t)
	h := hex.EncodeToString([]byte(t1))
	log.Print("Target UID:", h)
}

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
