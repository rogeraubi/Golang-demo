package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r:=mux.NewRouter();
	r.HandleFunc("/book/{title}/page/{page}",func ( w http.ResponseWriter, r *http.Request)  {
		vars := mux.Vars(r)
		fmt.Println(vars)
		title := vars["title"]
		page  := vars["page"]
		fmt.Fprint(w,"you have done %s: on the page %s\n",title,page)
		  fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
		
	}).Methods("GET") 
	http.ListenAndServe(":80",r)
}