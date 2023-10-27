package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/message", messageHandler)
    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Hello, Go!</title>
            <script src="https://unpkg.com/htmx.org@1.9.6"></script>
        </head>
        <body>
            <h1>Welcome to the go web server</h1>
            <button hx-get="/message" hx-target="#result">Fetch Message</button>
            <div id="result"></div>
        </body>
        </html>
    `)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "This is from the message route")
}
