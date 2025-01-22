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
            route.Handler(w, r)
        } else {
            http.NotFound(w, r)
        }
    })
}

func (server *Server) Start() error {
    server.handleIncomingRequests()
    err := http.ListenAndServe(fmt.Sprintf(":%d", server.port), nil)

    return err
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

func (app *Server) Post(path string, routeHandler func(http.ResponseWriter, *http.Request)) {
    app.router.AddRoute("POST", path, routeHandler)
}



