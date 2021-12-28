package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var p struct {
			Human struct {
				Age int `json:"age,omitempty"`
			} `json:"human,omitempty"`
		}

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if p.Human.Age > 0 {
			p.Human.Age += 10
		} else {
			http.Error(w, "no age set", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)

	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}

}
