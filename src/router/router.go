package router

import "net/http"

type Route struct {
    Method string
    Path string
    Handler func(w http.ResponseWriter,r *http.Request)
}

type Router struct {
    routeList []Route
}

func (r *Router) AddRoute(method string, path string, handler func(http.ResponseWriter, *http.Request)) {
    route := Route {
        method,
        path,
        handler,
    }

    r.routeList = append(r.routeList, route)
}

func (router *Router) FindRoute(method string, path string) *Route {
    for _, route := range router.routeList {
        if route.Method == method && route.Path == path {
            return &route
        }
    }

    return nil 
}

