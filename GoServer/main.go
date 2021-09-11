package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PostItem() {
	http.HandleFunc("/postItem", func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w, r)
		url := "https://todo-list-18b93-default-rtdb.firebaseio.com/items.json"
		fmt.Println("URL:> ", url)

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		fmt.Printf(string(body))

		var jsonStr = []byte(string(body))
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	})
}

func GetItems() {
	http.HandleFunc("/getItems", func(w http.ResponseWriter, r *http.Request) {
		url := "https://todo-list-18b93-default-rtdb.firebaseio.com/items.json"
		fmt.Println("URL:> ", url)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

}

func DeleteItem() {
	http.HandleFunc("/deleteItem", func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w, r)
		IDs, ok := r.URL.Query()["removeID"]

		if !ok || len(IDs[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}
		id := IDs[0]

		url := "https://todo-list-18b93-default-rtdb.firebaseio.com/items/" + id + ".json"
		fmt.Println("URL:> ", url)

		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("resp:> ", resp)
	})
}

func main() {

	PostItem()
	GetItems()
	DeleteItem()

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
