package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gretarst/go-rest-api/internal/comment"
)

// Stores Pointer to our services
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

type Response struct {
	Message string
}

// Returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// Setups the routes of the application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable parse comment id")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Error retrieving comment")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Failed to retrive all comments")
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode JSON Body")
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode JSON Body")
	}

	vars := mux.Vars(r)
	id := vars["id"]
	commendID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse comment id")
	}

	comment, err = h.Service.UpdateComment(uint(commendID), comment)
	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	commendID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse comment id")
	}

	err = h.Service.DeleteComment(uint(commendID))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete comment")
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Comment Deleted"}); err != nil {
		panic(err)
	}
}
