package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Stores Pointer to our services
type Handler struct {
	Router *mux.Router
}

// Returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// Setups the routes of the application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})

}
