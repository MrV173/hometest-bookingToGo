package routes

import (
	"test-gorilla-mux/handlers"
	"test-gorilla-mux/pkg/mysql"
	"test-gorilla-mux/repositories"

	"github.com/gorilla/mux"
)

func FamilyRoutes(r *mux.Router) {
	familyRepository := repositories.RepositoryFamily(mysql.DB)
	h := handlers.HandlerFamily(familyRepository)

	r.HandleFunc("/families", h.FindFamilies).Methods("GET")
	r.HandleFunc("/family/{id}", h.GetFamily).Methods("GET")
	r.HandleFunc("/family", h.CreateFamily).Methods("POST")
	r.HandleFunc("/family/{id}", h.UpdateFamily).Methods("PATCH")
	r.HandleFunc("/family/{id}", h.DeleteFamily).Methods("DELETE")
}
