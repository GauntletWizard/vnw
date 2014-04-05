package ui

import (
	"net/http"
  "flag"
  "vnw/config"
  "vnw/config"
  "html/template"
  "fmt"
  "log"
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


