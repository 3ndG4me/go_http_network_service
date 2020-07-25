package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var endpoint = "http://127.0.0.1:8080/"

func main() {
	do_get()
	do_post()
}

func do_get() {
	resp, err := http.Get(endpoint + "/TEST")

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

func do_post() {

	json_data := map[string]string{"item": "My Data"}
	payload, _ := json.Marshal(json_data)
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))

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
