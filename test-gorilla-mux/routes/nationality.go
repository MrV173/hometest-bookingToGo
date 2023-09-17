package routes

import (
	"test-gorilla-mux/handlers"
	"test-gorilla-mux/pkg/mysql"
	"test-gorilla-mux/repositories"

	"github.com/gorilla/mux"
)

func NationalityRoutes(r *mux.Router) {
	nationalityRepository := repositories.RepositoryNationality(mysql.DB)
	h := handlers.HandlerNationality(nationalityRepository)

	r.HandleFunc("/nationalities", h.FindNationalities).Methods("GET")
	r.HandleFunc("/nationality/{id}", h.GetNationality).Methods("GET")
	r.HandleFunc("/nationality", h.CreateNationality).Methods("POST")
	r.HandleFunc("/nationality/{id}", h.UpdateNationality).Methods("PATCH")
	r.HandleFunc("/nationality/{id}", h.DeleteNationality).Methods("DELETE")
}
