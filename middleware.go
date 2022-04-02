package main

import "net/http"

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
