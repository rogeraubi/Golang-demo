package main

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
