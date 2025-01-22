package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HomeRouteHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./src/html/index.html")
}

func GetJsonResponse(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "User created!")
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Unable to read request body", http.StatusBadRequest)
        return
    }

    var data map[string]interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "User:\n%+v\n", data)
    fmt.Printf("User:\n%+v\n", data)
}

