package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func RouteHandler() {
    http.HandleFunc("/", homePageRoute)
    http.HandleFunc("/mariana", marianaPageRoute)
}

func homePageRoute(w http.ResponseWriter, r *http.Request) {
    // http.Header.Set()
    file, err := os.Open("./src/html/index.html")
    if err != nil {
        log.Fatal(err)
        return
    }

    defer file.Close()

    fileInfo, err := file.Stat()

    if err != nil {
        log.Fatal(err)
        return
    }

    http.ServeContent(w, r, "../html/index.html", fileInfo.ModTime(), file)
}

func marianaPageRoute(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "A Mariana esteve aqui!\n")
}

