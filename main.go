package meal_reminder_server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aarrico/meal-reminder-server/meals"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	session    *mgo.session
	collection *mgo.Collection
}

type Supplement struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json: "name" bson:"name"`
	Amount float32       `json: "amount" bson:"amount"`
	Unit   string        `json: "unit" bson:"unit"`
	Days   []string      `json: "days" bson:"days"`
}

func (db *DB) GetSupplement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	var supp Supplement

	err := db.collection.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&supp)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(supp)
		w.Write(response)
	}
}

func (db *DB) PostSupplement(w http.ResponseWriter, r *http.Request) {
	var supp Supplement
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &supp)

	supp.ID = bson.NewObjectId()
	err := db.collection.Insert(supp)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(supp)
		w.Write(response)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	mealRouter = meals.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(mealRouter)))

	session, err := mgo.Dial("127.0.0.1")
	meals := session.DB("appdb").C("meals")
	supps := session.DB("appdb").C("supplements")
	mealDB := &DB{session: session, collection: meals}
	suppDB := &DB{session: session, collection: supps}

	if err != nil {
		panic(err)
	}
	defer session.Close()

	r := mux.NewRouter()
	r.HandleFunc("/v1/meals/{id:[a-zA-Z0-9]*}", mealDB.GetMeal).Methods("GET")
	r.HandleFunc("/v1/meals", mealDB.PostMeal).Methods("POST")
	r.HandleFunc("/v1/supps/{id:[a-zA-Z0-9]*}", suppDB.GetSupplement).Methods("GET")
	r.HandleFunc("/v1/supps", suppDB.PostSupplement)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
