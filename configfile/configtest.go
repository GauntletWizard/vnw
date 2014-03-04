package config

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var file = flag.String("dbfile", "foo.csv", "location to read/store the user database")
var reqpath = flag.String("reqpath", "http://tcbtech.org/~ted/foo.csv", "URL of member list")

const tmpext = ".tmp"

type Member struct {
	Name    string
	Id      int
	IdCards []string
}

type MemberList []Member

func Start() {
	log.Print("Opening file ", *file, " for config database")
	var c http.Client
	tmpfile := *file + tmpext
	for {
		resp, err := c.Get(*reqpath)
		if (err == nil) && (resp.StatusCode == 200) {
			// Write response to file
			f, err := os.Create(tmpfile)
			if err != nil {
				log.Fatal("Failed to open temp file ", tmpfile)
			}
			io.Copy(f, resp.Body)
			f.Close()

			// Validate, then populate
			_ = loadMembers(tmpfile)

		}
	}
}

func loadMembers(fname string) (l MemberList) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal("Failed to read from temp file ", fname)
	}
	csvReader := csv.NewReader(f)
	for read, err := csvReader.Read(); err == nil; read, err = csvReader.Read() {
		id, err := strconv.Atoi(read[1])
		if err != nil {
			log.Fatal("Bad Member ID ", read[1])
		}
		m := Member{Name: read[0],
			Id: id}
		m.IdCards = append([]string{}, read[2:]...)
		l = append(l, m)
	}
	return
}
