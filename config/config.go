package config

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var file = flag.String("dbfile", "foo.csv", "location to read/store the user database")
var reqpath = flag.String("reqpath", "http://tcbtech.org/~ted/stuff/foo.csv", "URL of member list")
var sleep = flag.Int("sleeptime", 10, "Number of seconds between updates of configfile")

const tmpext = ".tmp"

type Member struct {
	Name    string
	Id      int
	IdCards []string
}

type Cardlist map[string]*Member

var Cards *Cardlist

func init() {
	c :=make(Cardlist)
	Cards = &c
}

func Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Print("Opening file ", *file, " for config database")
	var c http.Client
	tmpfile := *file + tmpext
	cards := loadMembers(*file)
	Cards = &cards
	go func() {
		for {
			resp, err := c.Get(*reqpath)
			if (err == nil) && (resp.StatusCode == 200) {
				log.Print("Got config from server")
				// Write response to file
				f, err := os.Create(tmpfile)
				if err != nil {
					log.Fatal("Failed to open temp file ", tmpfile)
				}
				io.Copy(f, resp.Body)
				f.Close()

				// Validate, then populate
				list := loadMembers(tmpfile)
				if validateCardlist(&list) {
					Cards = &list
					os.Remove(*file)
					log.Print("Updating config file")
					os.Link(tmpfile, *file)
					os.Remove(tmpfile)
				} else {
					log.Print("Config failed to validate!")
				}

			} else {
				log.Print("Failed to get config from server: ", err, resp.StatusCode)
			}
			time.Sleep(time.Duration(*sleep) * time.Second)
		}
	}()
}

func loadMembers(fname string) (l Cardlist) {
	f, err := os.Open(fname)
	if err != nil {
		log.Println("Failed to read from file ", fname)
		return nil
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
		l[m.Name] = &m
	}
	return
}

func validateCardlist(l *Cardlist) bool {
	if (*l)["ERROR"] != nil {
		return false
	}
	jp := false
	dave := false
	ted := false
	// Check that the owners are in the DB
	for _, m := range *l {
		switch m.Name {
		case "Ted Hahn":
			ted = true
		case "JP Sugarbroad":
			jp = true
		case "David Stansel-Garner":
			dave = true
		}
	}
  if !(dave && ted && jp) {
    log.Print("D, T, JP:", dave, ted, jp)
  }
	return (dave && ted && jp)
}

func (m *Member) Log(id string) {
	log.Print("Member " + m.Name + " opened door with ID " + id)
}
