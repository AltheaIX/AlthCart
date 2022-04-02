package main

import (
	"encoding/base64"
	"net/http"
	"time"
)

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
