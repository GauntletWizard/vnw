package ui

import (
	"html/template"
	"log"
	"net/http"
	"vnw/core"
	"vnw/config"
)

var Httplistener string

func Start() {
	http.HandleFunc("/cards", listCards)
	http.HandleFunc("/status", statuss)
	http.HandleFunc("/testcard", testcard)
	http.HandleFunc("/clear", clearfailed)
	http.HandleFunc("/update", update)
	err := http.ListenAndServe(Httplistener, nil)
	if err != nil {
		log.Fatal("Failed to create server!", err)
	}
}

func listCards(w http.ResponseWriter, r *http.Request) {
  t, err := template.New("cardzs").Parse(`<!DOCTYPE HTML><html><head><title>List of member's cards.</title></head>
<body>
{{define "member"}}<div>{{.Name}}</div><div>{{.Id}}</div>{{end}}
{{range $c, $m := .}}<div>{{$c}}{{template "member" $m}}</div>{{end}}
</body></html>`)
/*	for i, j := range *config.Cards {
		w.Write([]byte(i))
		w.Write([]byte(fmt.Sprint(j)))
	} */
  if err != nil {
    log.Panic("Your template is bad, and you should feel bad!", err)
  }
  err = t.Execute(w, config.Cards)
  if err != nil {
    log.Panic("Template Execution failed.", err)
  }

}

func statuss(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("statuszs").Parse(`<!DOCTYPE HTML><html><head><title>Status Page for Lock Mechanism</title></head>
<body><div>Recent Failed{{range $i, $e := .}}<div>{{$i}}</div>{{end}}</div><div><form action="/clear"><input type="submit" value="Clear"></form></div>
<div><form action="/testcard">Test Card:<input name="cardid"><input type="submit"></form><form action="/update"><input type=submit value="Update"></form></body>`)
	if err != nil {
		log.Panic("Your template is bad, and you should feel bad!", err)
	}
	err = t.Execute(w, core.Failed)
	if err != nil {
		log.Panic("Template Execution failed.", err)
	}
}

func testcard(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for _, c := range r.Form["cardid"] {
		core.Auth <- c
	}
	w.Header().Add("Location", "/status")
	w.WriteHeader(303)
}

func clearfailed(w http.ResponseWriter, r *http.Request) {
	core.Clear()
	w.Header().Add("Location", "/status")
	w.WriteHeader(303)
}

func update(w http.ResponseWriter, r *http.Request) {
  config.Update()
	w.Header().Add("Location", "/status")
	w.WriteHeader(303)
}

