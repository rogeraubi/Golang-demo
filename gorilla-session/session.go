package main 
import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

var (
	key = []byte("super-secret-key")
    store = sessions.NewCookiesStore(key)
) 

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ = store.Get()
}
