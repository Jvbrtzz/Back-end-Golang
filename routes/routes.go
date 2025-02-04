package routes

import (
	"log"
	"net/http"

	"github.com/Jvbrtzz/Back-end-golang/controllers" // Importação correta
	"github.com/Jvbrtzz/Back-end-golang/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	PORT := ":8000"

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	// Usando o handler GetAllUsers
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUsersById).Methods("GET")

	r.HandleFunc("/cards/{id}", controllers.GetUserCard).Methods("GET")
	r.HandleFunc("/comment/{id}", controllers.GetUserCardComments).Methods("GET")

	log.Fatal(http.ListenAndServe(PORT, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
