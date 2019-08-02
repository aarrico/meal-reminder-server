package supplements

import (
	"fmt"
)

type Repository struct{}

const SERVER = "http://localhost:27017"
const DBNAME = "mealreminder"
const COLLECTION = "supplements"

func (r Repository) GetSupplement(id) Supplement {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := Supplement{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results: ", err)
	}

	return results
}
