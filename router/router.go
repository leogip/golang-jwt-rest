package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/leogip/golang-jwt-rest/handlers"
	"github.com/leogip/golang-jwt-rest/middleware"
)

// NewRouter ...
func NewRouter() *mux.Router {
	routes := mux.NewRouter()
	
	fs := http.FileServer(http.Dir("./static"))
	routes.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	routes.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	routes.HandleFunc("/refresh", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/refresh.html")
	})
	routes.HandleFunc("/remove", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/remove.html")
	})

	routes.HandleFunc("/api/users", handlers.GetAllUsers).Methods("GET")
	routes.HandleFunc("/api/tokens", handlers.GetAllTokens).Methods("GET")

	routes.HandleFunc("/api/token/get/{userId}", handlers.GetTokens).Methods("GET")
	routes.Handle("/api/token/refresh",
		middleware.AuthRequired(http.HandlerFunc(handlers.RefreshTokens))).Methods("GET")
		routes.Handle("/api/token/remove/t/{tokenId}",
		middleware.AuthRequired(http.HandlerFunc(handlers.RemoveToken))).Methods("DELETE")
	routes.Handle("/api/token/remove/u/{userId}",
		middleware.AuthRequired(http.HandlerFunc(handlers.RemoveAllTokens))).Methods("DELETE")

	return routes
}
