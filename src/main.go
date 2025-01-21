package main

import (
	"fmt"
	"go-api/src/core"
	"go-api/src/routes"
	"net/http"
)

func main() {
    const PORT int = 8080

    app := core.CreateServer(PORT)

    app.Get("/", routes.HomeRouteHandler)
    app.Get("/test", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "This is working!!!\n");
    })

    app.Start()
}

