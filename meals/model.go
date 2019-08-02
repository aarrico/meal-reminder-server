package meals

import "gopkg.in/mgo.v2/bson"

type Food struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

type Component struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Food   Food          `json:"food" bson:"food"`
	Amount float32       `json:"amount" bson:"amount"`
}

type Meal struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Order uint32        `json:"order" bson:"order"`
	Foods []Component   `json:"foods" bson:"foods"`
}

type Meals []Meal
