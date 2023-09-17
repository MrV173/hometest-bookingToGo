package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	CustomerRoutes(r)
	FamilyRoutes(r)
	NationalityRoutes(r)
}
