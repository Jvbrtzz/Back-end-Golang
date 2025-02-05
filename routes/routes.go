package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/Jvbrtzz/Back-end-golang/controllers"
	"github.com/Jvbrtzz/Back-end-golang/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func HandleRequest() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUsersById).Methods("GET")
	r.HandleFunc("/registeruser", controllers.RegisterNewUser).Methods("POST")

	r.HandleFunc("/cards/{id}", controllers.GetUserCard).Methods("GET")
	r.HandleFunc("/registercard", controllers.RegisterNewCard).Methods("POST")

	r.HandleFunc("/comment/{id}", controllers.GetUserCardComments).Methods("GET")

	log.Fatal(http.ListenAndServe(PORT, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
