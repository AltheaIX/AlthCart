package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Products struct {
	Id   int
	Name string
	Desc string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/althcart")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func productList() ([]Products, error) {
	data := []Products{}

	db, err := connect()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	res, err := db.Query("SELECT * FROM products")
	defer res.Close()
	if err != nil {
		return nil, err
	}

	for res.Next() {
		elem := Products{}
		err := res.Scan(&elem.Id, &elem.Name, &elem.Desc)
		if err != nil {
			return nil, err
		}
		data = append(data, elem)
	}
	return data, nil
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index.html").ParseFiles("index.html"))
	data, err := productList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerCart(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").ParseFiles("index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
