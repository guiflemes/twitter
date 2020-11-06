package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	_"github.com/guiflemes/twitter_clone/logger"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"github.com/guiflemes/twitter_clone/app/middleware"
)

type Route struct {
	Method  string
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var (
	routes     = make([]Route, 0)
)

//Register a list or single a single route
func RegisterRouter(rs ...Route) {
	for _, r := range rs{
		routes = append(routes, r)
	}
}

func Handlers() {

	router := mux.NewRouter()

	for _, rt := range routes {
		router.HandleFunc("/" + rt.Path, middleware.CheckDB(rt.Handler)).Methods(rt.Method)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", PORT), handler))
}
