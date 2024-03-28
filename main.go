package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status Status `json:"status"`
}

func main() {
	go func() {
		for {
			data := Data{
				Status: Status{
					Water: rand.Intn(100) + 1,
					Wind:  rand.Intn(100) + 1,
				},
			}

			file, _ := json.MarshalIndent(data, "", " ")

			_ = ioutil.WriteFile("status.json", file, 0644)

			time.Sleep(15 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "status.json")
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
