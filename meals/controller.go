package meals

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) GetMeal(w http.ResponseWriter, r *http.Request) {
	meal := c.Repository.GetMeal(id)

	data, _ := json.Marshal(meal)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	return
}

func (c *Controller) AddMeal(w http.ResponseWriter, r *http.Request) {
	var meal Meal
	body, err := ioutil.ReadAll(r.Body)

	log.Println(body)

	if err != nil {
		log.Fatalln("Error AddMeal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddMeal", err)
	}

	if err := json.Unmarshal(body, &meal); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddMeal unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(meal)
	success := c.Repository.AddMeal(meal)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
