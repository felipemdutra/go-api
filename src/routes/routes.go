package routes

import (
	"net/http"
)

func HomeRouteHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./src/html/index.html")
}

