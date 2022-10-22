package routes

import (
	"foodways/handlers"
	"foodways/pkg/mysql"
	"foodways/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	// r.HandleFunc("/register", middleware.UploadFile(h.Register)).Methods("POST")

	r.HandleFunc("/login", h.Login).Methods("POST")
}
