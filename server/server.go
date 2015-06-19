package server

import (
	"log"
	"net/http"
	"path"
	"strings"
)

// New creates new media server as http.Handler
func New(dir string) *Server {
	d := http.Dir(dir)
	return &Server{
		root: &d,
	}
}

// Server implements http.Handler
type Server struct {
	root *http.Dir
}

// ServeHTTP implements http.Handler method
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}

	log.Printf("access path: %s", r.URL.Path)
	log.Printf("access upath: %s", upath)
	log.Printf("access: %s", path.Join(string(*s.root), path.Clean(upath)))
	http.ServeFile(w, r, path.Join(string(*s.root), path.Clean(upath)))
}
