package main

import (
  "log"
  "net/http"
  "html/template"
  "os"
)

var tpl = template.Must(template.ParseFiles("static/main.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
  port := os.Getenv("PORT")
  if port == "" {
		port = "8080"
	}
		
  mux := http.NewServeMux()

  fs := http.FileServer(http.Dir("public"))
  mux.Handle("/assets/", http.StripPrefix("/public/", fs))

  mux.HandleFunc("/", indexHandler)

  log.Println("Listening on",port)
  http.ListenAndServe(":"+port, mux)
 }
