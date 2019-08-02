package supplements

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) GetSupplement(w http.ResponseWriter, r *http.Request) {
	supp := c.Repository.GetSupplement(id)

	data, _ := json.Marshal(supp)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	return
}

func (c *Controller) AddSupplement(w http.ResponseWriter, r *http.Request) {
	var supp Supplement
	body, err := ioutil.ReadAll(r.Body)

	log.Println(body)

	if err != nil {
		log.Fatalln("Error AddSupplement", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddSupplement", err)
	}

	if err := json.Unmarshal(body, &supp); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddSupplement unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(supp)
	success := c.Repository.AddSupplement(supp)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
