package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)
var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret ( w http.ResponseWriter, r *http.Request)  {

	session, _ := store.Get(r, "cookie")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden",http.StatusForbidden)
		return 
	}

	fmt.Fprintln(w,"The cake is a lie !!")
	
}

func login (w http.ResponseWriter, r *http.Request)  {
	session , _ := store.Get(r, "cookie")
	fmt.Printf("session%v", session)
	session.Values["authenticated"] = true
	session.Save(r,w)
	
}

func logout (w http.ResponseWriter, r *http.Request) {
	 session, _ :=store.Get(r,"cookie")
	 session.Values["authenticated"] = false
	 session.Save(r ,w)
}

func main() {
   http.HandleFunc("/secret",secret)
   http.HandleFunc("/login",login)
   http.HandleFunc("/logout",logout)
   http.ListenAndServe(":8080",nil)
}

/**
$ go run sessions.go

$ curl -s http://localhost:8080/secret
Forbidden

$ curl -s -I http://localhost:8080/login
Set-Cookie: cookie-name=MTQ4NzE5Mz...

$ curl -s --cookie "cookie-name=MTQ4NzE5Mz..." http://localhost:8080/secret
The cake is a lie!
**/