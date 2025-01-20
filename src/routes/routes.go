package routes

import (
	"log"
	"net/http"
	"os"
)

func RouteHandler() {
    http.HandleFunc("/", homePageRoute)
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

