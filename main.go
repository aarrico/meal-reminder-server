package meal_reminder_server

import (
	"github.com/aarrico/meal-reminder-server/meals"
	"github.com/aarrico/meal-reminder-server/supplements"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

type DB struct {
	session    *mgo.session
	collection *mgo.Collection
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	mealRouter = meals.NewRouter()
	supplementRouter = supplements.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(mealRouter, supplementRouter)))
}
