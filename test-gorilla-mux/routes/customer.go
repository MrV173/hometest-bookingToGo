package routes

import (
	"test-gorilla-mux/handlers"
	"test-gorilla-mux/pkg/mysql"
	"test-gorilla-mux/repositories"

	"github.com/gorilla/mux"
)

func CustomerRoutes(r *mux.Router) {
	customerRepository := repositories.RepositoryCustomer(mysql.DB)
	h := handlers.HandlerCustomer(customerRepository)

	r.HandleFunc("/customers", h.FindCustomers).Methods("GET")
	r.HandleFunc("/customer/{id}", h.GetCustomer).Methods("GET")
	r.HandleFunc("/customer", h.CreateCustomer).Methods("POST")
	r.HandleFunc("/customer/{id}", h.UpdateCustomer).Methods("PATCH")
	r.HandleFunc("/customer/{id}", h.DeleteCustomer).Methods("DELETE")
}
