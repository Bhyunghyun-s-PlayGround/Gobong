package main

import(
	"encoding/json"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Router struct {
	routes map[string]map[string]HandlerFunc
}

func NewRouter() * Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (r *Router) Handle(method, path string, handler HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]HandlerFunc)
	}
	r.routes[path][method] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if methods, ok := r.routes[req.URL.Path]; ok {
		if handler, ok := methods[req.Method]; ok {
			handler(w, req)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, req)
}

func JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}