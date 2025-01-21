package core

import (
	"fmt"
	"net/http"
    "go-api/src/router"
)

type Server struct {
    port int
    router router.Router
}

func (s *Server) handleIncomingRequests() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        route := s.router.FindRoute(r.Method, r.URL.Path)
        if route != nil {
            fmt.Printf("New %s connection from %s\n", r.Method, r.URL.Path)
            route.Handler(w, r)
        } else {
            http.NotFound(w, r)
        }
    })
}

func (server *Server) Start() {
    server.handleIncomingRequests()
    http.ListenAndServe(fmt.Sprintf(":%d", server.port), nil)
}

func CreateServer(port int) *Server {
    var server *Server = &Server {
        port: port,
        router: router.Router{},
    }

    return server
}

func (app *Server) Get(path string, routeHandler func(http.ResponseWriter, *http.Request)) {
    app.router.AddRoute("GET", path, routeHandler)

}

