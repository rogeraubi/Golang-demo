package main
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)
var (
	count    map[string]int
	countMux sync.Mutex
	store    *sessions.CookieStore
)

type User struct {
	Username string
	Password string
}

func init() {
	count = make(map[string]int)
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	username := session.Values["username"]
	if username != nil {
		countMux.Lock()
		userCount := count[username.(string)]
		countMux.Unlock()
		tmpl := template.Must(template.ParseFiles("static/counter.html"))
		tmpl.Execute(w, userCount)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/login.html"))
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		user := User{Username: username, Password: password}

		session, _ := store.Get(r, "session-name")
		session.Values["username"] = user.Username
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		tmpl.Execute(w, nil)
	}
}
func incrementHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	username := session.Values["username"]
	if username == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	countMux.Lock()
	count[username.(string)]++
	userCount := count[username.(string)]
	countMux.Unlock()
	fmt.Fprintf(w, strconv.Itoa(userCount))
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/increment", incrementHandler)
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
