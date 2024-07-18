package web

import (
	"log"
	"net/http"
	"sync"

	"github.com/nicolasparada/go-mux"
)

type Handler struct {
	Logger  *log.Logger
	Service *aboba.Service
	once    sync.Once
	handler http.Handler
}

func (h *Handler) init() {
	r := mux.NewRouter()

	r.Handle("/login", mux.MethodHandler{
		http.MethodGet:  h.showLogin,
		http.MethodPost: h.login,
	})

	h.handler = r
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.once.Do(h.init)
	h.handler.ServeHTTP(w, r)
}
