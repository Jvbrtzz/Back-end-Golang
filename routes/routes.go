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

	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusNoContent)
	})

	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUsersById).Methods("GET")
	r.HandleFunc("/registeruser", controllers.RegisterNewUser).Methods("POST")
	r.HandleFunc("/userLogin", controllers.UserLogin).Methods("POST")

	r.HandleFunc("/getUserCard/{id}", controllers.GetUserCard).Methods("GET")
	r.HandleFunc("/sharedCards/{id}", controllers.GetSharedCardsForUser).Methods("GET")
	r.HandleFunc("/sharedUsers/{id}", controllers.GetShareUserCard).Methods("GET")
	r.HandleFunc("/registercard", controllers.RegisterNewCard).Methods("POST")

	r.HandleFunc("/comment/{id}", controllers.GetCommentsByCard).Methods("GET")

	// Middleware de CORS aplicado corretamente
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	log.Printf("Servidor rodando na porta %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, corsMiddleware(r)))
}
