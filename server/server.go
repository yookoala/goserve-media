package server

import (
	"net/http"
)

func New(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}
