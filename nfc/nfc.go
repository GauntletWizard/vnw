package nfc

import (
	"encoding/hex"
	"flag"
	"log"
	"time"
	"unsafe"
	"vnw/core"
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

func Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cs := C.CString(*nfcdevice)
	C.init_globals(cs)
	C.free(unsafe.Pointer(cs))
}

func GetIds() []string {
	ids := make([]string, 0)
	nm := C.nfc_modulation{
		nmt: C.NMT_ISO14443A,
		nbr: C.NBR_106,
	}
	C.read_ids(nm)
	id := C.get_id()
	for id != nil {
		//		log.Print(ids)
		ids = append(ids, hex.EncodeToString([]byte(C.GoString(id))))
		C.free(unsafe.Pointer(id))
		id = C.get_id()
	}
	return ids
}
func Poll() {
	for {
		i := GetIds()
		//log.Print(i)
		if len(i) > 0 {
			for j := range i {
				core.Auth <- i[j]
			}
			time.Sleep(time.Millisecond * time.Duration(*nfcwait))
		} else {
			time.Sleep(time.Millisecond * time.Duration(*nfcpoll))
		}
	}
}
