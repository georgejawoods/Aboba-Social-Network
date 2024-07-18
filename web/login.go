package web

import (
	"aboba"
	"net/http"
	"net/url"
)

var loginTmpl = parseTmpl("login.tmpl")

type loginData struct {
	Form url.Values
	Err  error
}

func (h *Handler) renderLogin(w http.ResponseWriter, data loginData, statusCode int) {
	h.renderTmpl(w, loginTmpl, data, statusCode)
}

func (h *Handler) showLogin(w http.ResponseWriter, r *http.Request) {
	h.renderLogin(w, loginData{}, http.StatusOK)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	input := aboba.LoginInput{
		Email:    r.PostFormValue("email"),
		Username: formPtr(r.PostForm, "username"),
	}
	h.Service.Login(ctx, input)
}

func formPtr(form url.Values, key string) *string {
	if !form.Has(key) {
		return nil
	}

	s := form.Get(key)
	return &s
}
