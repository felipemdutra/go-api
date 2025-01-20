package core 

import (
	"fmt"
	"net/http"
)

type Server struct {
    port int
}

func (server *Server) Start() {
    http.ListenAndServe(fmt.Sprintf(":%d", server.port), nil)
}

func CreateServer(port int) *Server {
    var server *Server = &Server {
        port: port,
    }

    return server
}

func (app *Server) Get(route string, routeHandler func(http.ResponseWriter, *http.Request)) {
    http.HandleFunc(route, routeHandler)
}

