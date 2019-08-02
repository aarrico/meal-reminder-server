package supplements

import "gopkg.in/mgo.v2/bson"

type Supplement struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json: "name" bson:"name"`
	Amount float32       `json: "amount" bson:"amount"`
	Unit   string        `json: "unit" bson:"unit"`
	Days   []string      `json: "days" bson:"days"`
}

type Supplements []Supplement
