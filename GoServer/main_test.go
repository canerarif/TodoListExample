package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPostItem(t *testing.T) {
	url := "https://todo-list-18b93-default-rtdb.firebaseio.com/items.json"

	var jsonStr = []byte("{\"item\":\"merhaba\"}")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Item Post Error")
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Status Code Error")
	}
}

func TestGetItems(t *testing.T) {
	url := "https://todo-list-18b93-default-rtdb.firebaseio.com/items.json"

	resp, err := http.Get(url)
	if err != nil {
		t.Error("Items Get Error")
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Get Item Status Code Error")
	}
}

func TestDeleteItems(t *testing.T) {
	url := "https://todo-list-18b93-default-rtdb.firebaseio.com/items/1.json"

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Error("Delete Item Error")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Error("Delete Item Status Code Error")
	}
}
