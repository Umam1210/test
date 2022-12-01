package routes

import (
	"journey/handlers"
	"journey/pkg/middleware"
	"journey/pkg/mysql"
	"journey/repositories"

	"github.com/gorilla/mux"
)

func ArtikelRoutes(r *mux.Router) {
	artikelRepository := repositories.RepositoryArtikel(mysql.DB)
	h := handlers.HandlerEpresence(artikelRepository)

	r.HandleFunc("/epresences", h.FindEpresences).Methods("GET")
	r.HandleFunc("/epresences/user/{user_id}", h.FindEpresencesbyUserId).Methods("GET")
	r.HandleFunc("/epresence/{id}", h.GetEpresence).Methods("GET")
	r.HandleFunc("/epresence", middleware.Auth(h.CreateEpresence)).Methods("POST")
	r.HandleFunc("/epresence/{id}", h.UpdateEpresence).Methods("PATCH")
	r.HandleFunc("/epresence/{id}", h.DeleteEpresence).Methods("DELETE")

}
