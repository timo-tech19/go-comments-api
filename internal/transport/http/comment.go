package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/timo-tech19/go-comments-api/internal/comment"
)

// Defines functionality provided by our comment service.
type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

type Response struct {
	Message string
}

// Defines data and validators for expected data in request body.
type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

// Converts PostCommentResquest struct data expected from the request into Comment struct data used in the comment service
func postCommentRequestToComment(c PostCommentRequest) comment.Comment {
	// the ID field is set to it's zero value since it is not initialized here.
	return comment.Comment{
		Slug:   c.Slug,
		Author: c.Author,
		Body:   c.Body,
	}
}

// Add a new comment to comment service.
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var postCmt PostCommentRequest

	// decode request body and scan into postCmt
	if err := json.NewDecoder(r.Body).Decode(&postCmt); err != nil {
		return
	}

	// validate postCmt using field tag validators
	validate := validator.New()
	err := validate.Struct(postCmt)

	if err != nil {
		http.Error(w, "Not a valid Comment", http.StatusBadRequest)
		return
	}

	cmt := postCommentRequestToComment(postCmt)

	postedCmt, err := h.Service.PostComment(r.Context(), cmt)
	if err != nil {
		log.Print(err)
		return
	}

	// Send json response
	if err := json.NewEncoder(w).Encode(postedCmt); err != nil {
		panic(err)
	}
}

// Retrieve comment by id(specified in url param) from comment service
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

// Updates comment by id(specified in url param) in comment service
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var cmt comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

// Delete comment by id(specified in url param) in comment service
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID := vars["id"]
	if commentID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), commentID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Sucessfully deleted"}); err != nil {
		panic(err)
	}
}
