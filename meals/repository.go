package meals

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type Repository struct{}

const SERVER = "http://localhost:27017"
const DBNAME = "mealreminder"
const COLLECTION = "meals"

func (r Repository) GetMeal(id) Meal {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establist connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := Meal{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results: ", err)
	}

	return results
}
