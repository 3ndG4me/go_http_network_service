package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// A URL and url you want to send data to
var url = "http://127.0.0.1:8081/"
var endpoint = "test"

func main() {
	http.HandleFunc("/", service_handler)
	http.ListenAndServe(":8080", nil)
}

func service_handler(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path[1:]

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	data := string(body)

	switch req.Method {
	case "GET":
		if path != "" {
			log.Println("GET: " + path)
			fmt.Fprintf(w, "%s", path)
			url_get_request(path)
		}
	case "POST":
		if data != "" {
			log.Println("POST: " + data)
			fmt.Fprintf(w, "%s", data)
			url_post_request(data)
		}
	default:
		fmt.Println("Invalid Request Type")
	}

}

func url_get_request(path string) {
	resp, err := http.Get(url + path)

	if err != nil {
		log.Println("ERROR: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	body_str := string(body)

	if err != nil {
		log.Println("ERROR: %s", err)
	}

	fmt.Println(body_str)

}

func url_post_request(data string) {

	payload, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		log.Println("ERROR: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	body_str := string(body)

	if err != nil {
		log.Println("ERROR: %s", err)
	}
	if body_str != "" {
		fmt.Println(body_str)
	}
}
