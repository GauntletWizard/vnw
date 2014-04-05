package ui

import (
	"net/http"
  "flag"
  "vnw/config"
  "html/template"
)

var httplistener = flag.String("port", ":80", "Listen Address for webserver")

func init() {
  flag.Parse()
}

func Start() {
  http.HandleFunc("/cards", listCards)
  http.HandleFunc("/status", statuss)
  http.ListenAndServe(*httplistener, nil)
}

func listCards(w http.ResponseWriter, r *http.Request) {
  for i, j := range *config.Cards {
    w.Write(i, j)
  }
}

func statuss(w http.ResponseWriter, r *http.Request) {
  t, err := template.New("statuszs").Parse("<!DOCTYPE HTML><html><head><title>Status Page for Lock Mechanism")
  if err != nil {
    log.Error("Your template is bad, and you should feel bad!", err)
  }
}


