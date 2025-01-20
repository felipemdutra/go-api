package core 

import (
	"fmt"
	"go-api/src/routes"
	"net/http"
)

type Server struct {
    port int
}

func CreateServer(port int) *Server {
    return &Server {
        port: port,
    }
}

func StartServer(server *Server) {
    routes.RouteHandler()

    http.ListenAndServe(fmt.Sprintf(":%d", server.port), nil)
    fmt.Printf("Listening on port %d...", server.port)
}

