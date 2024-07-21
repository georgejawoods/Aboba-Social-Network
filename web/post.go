package web

import (
	"aboba"
	"net/http"
	"net/url"

	"github.com/nicolasparada/go-mux"
)

var postTmpl = parseTmpl("post.tmpl")

type postData struct {
	Session
	Post              aboba.PostRow
	Comments          []aboba.CommentsRow
	CreateCommentForm url.Values
	CreateCommentErr  error
}

func (h *Handler) renderPost(w http.ResponseWriter, data postData, statusCode int) {
	h.renderTmpl(w, postTmpl, data, statusCode)
}

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

// showPost handles GET /p/{postID}
func (h *Handler) showPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := mux.URLParam(ctx, "postID")
	p, err := h.Service.Post(ctx, postID)
	if err != nil {
		h.log(err)
		h.renderErr(w, r, err)
		return
	}

	cc, err := h.Service.Comments(ctx, postID)
	if err != nil {
		h.log(err)
		h.renderErr(w, r, err)
		return
	}

	h.renderPost(w, postData{
		Session:           h.sessionFromReq(r),
		Post:              p,
		Comments:          cc,
		CreateCommentForm: h.popForm(r, "create_comment_form"),
		CreateCommentErr:  h.popErr(r, "create_comment_err"),
	}, http.StatusOK)
}

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.putErr(r, "create_comment_err", err)
		http.Redirect(w, r, r.Referer(), http.StatusFound)
		return
	}

	ctx := r.Context()
	_, err := h.Service.CreateComment(ctx, aboba.CreateCommentInput{
		PostID:  r.PostFormValue("post_id"),
		Content: r.PostFormValue("content"),
	})

	if err != nil {
		h.putErr(r, "create_comment_err", err)
		h.session.Put(r, "create_comment_form", r.PostForm)
		http.Redirect(w, r, r.Referer(), http.StatusFound)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
