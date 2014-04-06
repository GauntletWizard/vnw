package ui

import (
	"net/http"
  "vnw/config"
  "vnw/core"
  "html/template"
  "fmt"
  "log"
)

var Httplistener string

func Start() {
  http.HandleFunc("/cards", listCards)
  http.HandleFunc("/status", statuss)
  err := http.ListenAndServe(Httplistener, nil)
  if err != nil {
    log.Fatal("Failed to create server!", err)
  }
}

func listCards(w http.ResponseWriter, r *http.Request) {
  for i, j := range *config.Cards {
    w.Write([]byte(i))
    w.Write([]byte(fmt.Sprint(j)))
  }
}

func statuss(w http.ResponseWriter, r *http.Request) {
  t, err := template.New("statuszs").Parse(`<!DOCTYPE HTML><html><head><title>Status Page for Lock Mechanism</title></head>
<body><div>Recent Failed{{range .}}<div>{{.}}</div>{{end}}</div></body>`)
  if err != nil {
    log.Panic("Your template is bad, and you should feel bad!", err)
  }
  t.Execute(w, core.Failed)
}


