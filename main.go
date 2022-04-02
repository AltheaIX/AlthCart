package main

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

func GetOnly(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		return false
	}
	return true
}

func Auth(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("Authentification")
	if err != nil {
		return false
	}
	return true
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

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	if !GetOnly(w, r) {
		http.Error(w, "Invalid method, only GET is allowed.", http.StatusBadRequest)
		return
	}

	if !Auth(w, r) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	c, err := r.Cookie("Authentification")

	baseUsername := c.Value
	tmpl := template.Must(template.New("index.html").ParseFiles("template/index.html"))
	dataProduct, err := productList()
	dataUserCart, err := cartData(baseUsername)

	userCartQuantity := getCartQuantity(dataUserCart)
	data := &IndexData{dataProduct, userCartQuantity}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerApiAdd(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("Authentification")
	if c.Value != "" && r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id := r.PostForm.Get("id")
		username := AuthToUsername(c.Value)
		db, err := connect()
		defer db.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		checkPrepare, err := db.Prepare("SELECT quantity from users_cart WHERE product_id=? AND username=?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		checkQuery, err := checkPrepare.Query(id, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer checkQuery.Close()

		if checkQuery.Next() {
			updateQuery, _ := db.Prepare("UPDATE users_cart SET quantity=quantity + 1 WHERE product_id=? AND username=?")
			updateQuery.Query(id, username)
			defer updateQuery.Close()
			return
		}

		createQuery, err := db.Prepare("INSERT INTO users_cart (product_id, username) VALUES (?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		createQuery.Query(id, username)
		defer createQuery.Close()
	}
}

func HandlerApiSetCookie(w http.ResponseWriter, r *http.Request) {
	cookieName := "Authentification"

	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c.Name = cookieName
		c.Path = "/"
		c.Value = base64.StdEncoding.EncodeToString([]byte("Malik"))
		c.Expires = time.Now().Add(1 * time.Hour)
		http.SetCookie(w, c)
	}
}

func HandlerCart(w http.ResponseWriter, r *http.Request) {
	if !GetOnly(w, r) {
		return
	}

	tmpl := template.Must(template.New("index").ParseFiles("template/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
