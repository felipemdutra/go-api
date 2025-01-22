package main

import (
	"go-api/src/core"
	"go-api/src/routes"
)

func main() {
    const PORT int = 8080

    app := core.CreateServer(PORT)

    app.Get("/", routes.HomeRouteHandler)
    app.Post("/create-user", routes.GetJsonResponse)

    app.Start()
}

