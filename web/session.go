package web

import (
	"aboba"
	"net/http"
)

type Session struct {
	IsLoggedIn bool
	User       aboba.User
}

func (h *Handler) sessionFromReq(r *http.Request) Session {
	var out Session

	if h.session.Exists(r, "user") {
		user, ok := h.session.Get(r, "user").(aboba.User)
		out.IsLoggedIn = ok
		out.User = user
	}

	return out
}
