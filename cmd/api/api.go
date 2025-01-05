package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/whatdislol/mobility-app/middleware"
	"github.com/whatdislol/mobility-app/service/stop"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		address: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	corsRouter := middleware.EnableCors(router)
    //router.Handle("/", http.FileServer(http.Dir("../static")))

	stopStore := stop.NewStore(s.db)
	stopHandler := stop.NewHandler(stopStore)
	stopHandler.RegisterRoutes(router)

	log.Println("Listening on", s.address)

	return http.ListenAndServe(s.address, corsRouter)
}