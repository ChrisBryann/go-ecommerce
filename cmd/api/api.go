package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ChrisBryann/go-ecommerce/services/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	usersStore := users.NewStore(s.db)
	usersHandler := users.NewHandler(usersStore)
	usersHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

func (s *APIServer) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
