package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8081", nil)
}

func HelloServer(w http.ResponseWriter, req *http.Request) {

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
		}
	case "POST":
		if data != "" {
			log.Println("POST: " + data)
			fmt.Fprintf(w, "%s", data)
		}
	default:
		fmt.Println("Invalid Request Type")
	}
}
