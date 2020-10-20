package handlers

//https://stackoverflow.com/questions/19356619/how-do-i-split-urls-in-multiple-files-using-gorilla-mux
import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/guiflemes/twitter_clone/logger"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

type Route struct {
	Method  string
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var (
	ApiVersion = "v2"
	routes = make([]Route, 0)
)

func RegisterRouter(r Route){
	routes = append(routes, r)
}

func Handlers() {
	router := mux.NewRouter()

	//router.HandleFunc("/register", middlew.CheckDB()).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	logger.Info("Application is about to begin")

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", PORT), handler))
}
