package main

import (
	"html/template"
	"net/http"
)

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
