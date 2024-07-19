package web

import (
	"aboba"
	"net/http"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.putErr(r, "create_post_err", err)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ctx := r.Context()
	_, err := h.Service.CreatePost(ctx, aboba.CreatePostInput{
		Content: r.PostFormValue("content"),
	})
	if err != nil {
		h.putErr(r, "create_post_err", err)
		h.session.Put(r, "created_post_form", r.PostForm)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
