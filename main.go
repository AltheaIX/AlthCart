package main

import (
	"net/http"
	"text/template"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("template/index.html").ParseFiles("template/index.html"))
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", HandlerIndex)
	mux.HandleFunc("/cart", HandlerCart)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("dependency"))))

	server := new(http.Server)
	server.Addr = ":8000"
	server.ListenAndServe()
}
