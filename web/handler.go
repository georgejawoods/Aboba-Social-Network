package web

import (
	"log"
	"net/http"
	"sync"

	"github.com/nicolasparada/go-mux"
)

type Handler struct {
	Logger  *log.Logger
	once    sync.Once
	handler http.Handler
}

func (h *Handler) init() {
	r := mux.NewRouter()

	r.Handle("/login", mux.MethodHandler{
		http.MethodGet: h.showLogin,
	})

	h.handler = r
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.once.Do(h.init)
	h.handler.ServeHTTP(w, r)
}
