package server

import (
	"docker-alarms/api/server/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

// Devolvemos un puntero con nuestro server
type Server struct {
	server *http.Server
}

// Inicializamos el servidor y montamos los endpoints
func New(port string) (*Server, error) {
	// Estructura que funciona de mux
	r := chi.NewRouter()

	// Se monta como raiz la direccion "api"
	r.Mount("/api/docker-firewatch", routes.New())

	serv := &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		ReadTimeout:       100 * time.Second,
		WriteTimeout:      100 * time.Second,
		TLSConfig:         nil,
		ReadHeaderTimeout: 100 * time.Second,
		IdleTimeout:       100 * time.Second,
		MaxHeaderBytes:    1000,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	// Construimos un server inicializado con el que acabamos de crear
	server := Server{server: serv}
	return &server, nil
}

func (serv *Server) Start() {
	log.Printf("Server running on port " + os.Getenv("PORT"))
	log.Fatal(serv.server.ListenAndServe())
}
