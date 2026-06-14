package routes

import (
	"net/http"

	"UniCore/internals/handlers"
)

func AuthRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/login", handlers.LoginUser)
	mux.HandleFunc("POST /user/register", handlers.RegisterUser)
	mux.HandleFunc("POST /user/refresh", handlers.RefreshHandler)
	mux.HandleFunc("POST /user/logout", handlers.LogoutHandler)

	return mux
}
