package nfc

import (
	//"vnw/config"
	"flag"
	"github.com/fuzxxl/nfc/latest/nfc"
	"log"
	// "time"
	"encoding/hex"
	"unsafe"
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
		log.Print(ids)
		ids = append(ids, hex.EncodeToString([]byte(C.GoString(id))))
		C.free(unsafe.Pointer(id))
		id = C.get_id()
	}
	return ids
}

func GetIdsold() []string {
	ids := make([]string, 0)
	count := C.int(0)
	// Also need to free the strings pointed to by t1, which we do below.
	t := C.get_ids(&count)
	defer C.free(unsafe.Pointer(t))
	log.Print(count, t)
	for i := 0; i < int(count); i++ {
		t1 := (*C.char)(unsafe.Pointer((uintptr(unsafe.Pointer(*t)) + (unsafe.Sizeof(t) * uintptr(i)))))
		h := hex.EncodeToString([]byte(C.GoString(t1)))
		log.Print("Target UID:", h)
		ids = append(ids, h)
		C.free(unsafe.Pointer(t1))
	}
	return ids
}

func Poll() {
	log.Print(GetIds())
}

func PollOld() {
	d, err := nfc.Open(*nfcdevice)
	if err != nil {
		log.Fatal("Failed to open NFC Device ", *nfcdevice, err)
	}
	nm := nfc.Modulation{
		Type:     nfc.ISO14443A,
		BaudRate: nfc.NBR_106,
	}

	targets, err := d.InitiatorListPassiveTargets(nm)
	log.Print(targets)
}
