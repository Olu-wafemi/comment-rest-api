package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// stores poniter to comment service
type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {

	return &Handler{}
}

// SetupRoutes- sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("api/health", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Effemm here")
	})

}
