package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

var INTERNAL_SERVER_ERROR = []byte("500: Internal Server Error")
var BAD_REQUEST = []byte("400: Bad Request")
var ERR_ALREADY_COMMITTED = "already been committed"

func New() http.Handler {
	r := chi.NewRouter()

	return r
}
