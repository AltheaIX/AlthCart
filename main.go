package main

import (
	"database/sql"
	"encoding/base64"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Products struct {
	Id   int
	Name string
	Desc string
}

type UserCart struct {
	Id        int
	ProductID int
	Username  string
	Quantity  int
}

type IndexData struct {
	Products  []Products
	CartCount int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/althcart")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getCartQuantity(data []UserCart) int {
	quantity := 0
	for _, each := range data {
		quantity += each.Quantity
	}
	return quantity
}

func cartData(base string) ([]UserCart, error) {
	data := []UserCart{}

	db, err := connect()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	username := AuthToUsername(base)
	checkPrepare, err := db.Prepare("SELECT * from users_cart WHERE username=?")
	checkQuery, err := checkPrepare.Query(username)
	defer checkQuery.Close()
	if err != nil {
		return nil, err
	}

	for checkQuery.Next() {
		elem := UserCart{}
		err := checkQuery.Scan(&elem.Id, &elem.ProductID, &elem.Username, &elem.Quantity)
		if err != nil {
			return nil, err
		}
		data = append(data, elem)
	}
	return data, nil
}

func productList() ([]Products, error) {
	data := []Products{}

	db, err := connect()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	res, err := db.Query("SELECT * FROM products LIMIT 10")
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

func AuthToUsername(base string) string {
	if base != "" {
		base64Decode, _ := base64.StdEncoding.DecodeString(base)
		username := string(base64Decode)
		return username
	}
	return ""
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", HandlerIndex)
	mux.HandleFunc("/login", HandlerCart)
	mux.HandleFunc("/cart", HandlerCart)
	mux.HandleFunc("/api/add", HandlerApiAdd)
	mux.HandleFunc("/api/setcookie", HandlerApiSetCookie)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("dependency"))))

	server := new(http.Server)
	server.Addr = ":8000"
	server.ListenAndServe()
}
