package config

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	File       string
	Reqpath    string
	Sleep      int
	Secret     []byte
	sa         smtp.Auth
	SAFile     string
	SMTPServer string
	Mailto     string
	mailto     []string
)

//= flag.Int("sleeptime", 600, "Number of seconds between updates of configfile")

const tmpext = ".tmp"

type Member struct {
	Name    string
	Id      int
	IdCards []string
}

type Cardlist map[string]*Member

var Cards *Cardlist
var update chan time.Time

func init() {
	c := make(Cardlist)
	Cards = &c
	update = make(chan time.Time, 0)
}

func Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Print("Opening file ", File, " for config database")
	var c http.Client
	tmpfile := File + tmpext
	cards := loadMembers(File)
	Cards = &cards
	timer := time.Tick(time.Duration(Sleep) * time.Second)
	pass, err := ioutil.ReadFile(SAFile)
	if err != nil {
		log.Print("No auth, not e-mailing")
	} else {
		hostname := strings.Split(SMTPServer, ":")
		sa = smtp.PlainAuth("", "ted@verneandwells.com", string(pass), hostname[0])
		mailto = strings.Split(Mailto, ",")
	}
	go func() {
		for {
			resp, err := c.Post(Reqpath, "text/plain", bytes.NewReader(Secret))
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
					os.Remove(File)
					log.Print("Updating config file")
					os.Link(tmpfile, File)
					os.Remove(tmpfile)
				} else {
					log.Print("Config failed to validate!")
				}

			} else if err != nil {
				log.Print("Failed to get config from server: ", err)
			} else {
				log.Print("Received error reading config from server: ", resp.StatusCode)
			}
			select {
			case <-timer:
			case <-update:
				log.Print("Update requested!")
			}
		}
	}()
}

func Update() {
	update <- time.Now()
}

func loadMembers(fname string) (l Cardlist) {
	log.Print("Opening Database: ", fname)
	f, err := os.Open(fname)
	if err != nil {
		log.Print("Failed to read from file ", fname)
		return nil
	}
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	csvReader.LazyQuotes = true
	csvReader.TrimLeadingSpace = true
	// log.Print("CSV Reader:", csvReader)
	read, err := csvReader.Read()
	log.Print(err)
	l = make(Cardlist, 100)
	for ; err == nil; read, err = csvReader.Read() {
		if err != nil {
			log.Fatal("Failed to read record:", err)
		}
		//log.Print(read)
		id, err := strconv.Atoi(read[1])
		if err != nil {
			log.Fatal("Bad Member ID ", read[1])
		}
		m := Member{Name: read[0],
			Id: id}
		m.IdCards = append([]string{}, read[2:]...)
		for _, i := range m.IdCards {
			l[i] = &m
		}
	}
	log.Print(err)
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
		//log.Print("Member:", m)
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
	lm := "Member " + m.Name + " opened door with ID " + id + " at " + time.Now().Format("Jan 2, 2006 at 3:04pm")
	log.Print(lm)
	if sa != nil {
		text := "From: ted@verneandwells.com\nto: david@verneandwells.com\nSubject: Door opened\n\n" + lm
		err := smtp.SendMail(SMTPServer, sa, "ted@verneandwells.com", mailto, []byte(text))
		if err != nil {
			log.Print("Failed to send mail!", err)
		}
	}
}
