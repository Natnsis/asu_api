package routes

import (
	"net/http"

	"UniCore/internals/handlers"
)

func AuthRouter(mux *http.ServeMux) {
	mux.HandleFunc("POST /user/register", handlers.RegisterUser)
	mux.HandleFunc("POST /user/login", handlers.LoginUser)
}
