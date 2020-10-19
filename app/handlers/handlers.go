package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/guiflemes/twitter_clone/logger"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handlers(){
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)


	logger.Info("Application is about to begin")

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", PORT), handler))
}