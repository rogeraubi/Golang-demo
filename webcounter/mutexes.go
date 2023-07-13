/*package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
)

var (
    count    int
    countMux sync.Mutex
)

func incrementHandler(w http.ResponseWriter, r *http.Request) {
    countMux.Lock()
    count++
    updatedCount := count
    countMux.Unlock()

    // Send the updated count as the response
    fmt.Fprintf(w, strconv.Itoa(updatedCount))
}

func main() {
    // Serve the index.html file
    http.Handle("/", http.FileServer(http.Dir(".")))
    

    // Define the /increment endpoint
    http.HandleFunc("/increment", incrementHandler)

    // Start the HTTP server
    log.Println("Server listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
*/

/*package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}
*/



package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))
var (
    count    int
    countMux sync.Mutex
)

func incrementHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    countMux.Lock()
    count++
    updatedCount := count
    countMux.Unlock()
    session.Values["count"] = updatedCount
    // Send the updated count as the response
    session.Save(r, w)
    fmt.Fprintf(w, strconv.Itoa(updatedCount))
}

func main() {
    // Serve the index.html file
    http.Handle("/", http.FileServer(http.Dir(".")))
    // Define the /increment endpoint
    http.HandleFunc("/increment", incrementHandler)
    // wrappedHandler := sessions.Middleware(store, nil)(http.DefaultServeMux)

    // Start the HTTP server
    log.Println("Server listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}