package handler

import (
	"around/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request){
	//1. process http request: json string -> post struct
	fmt.Println("Received one upload request")
    decoder := json.NewDecoder(r.Body)
    var p model.Post
    if err := decoder.Decode(&p); err != nil {
        panic(err)
    }
	fmt.Fprintf(w, "Post received: %s\n", p.Message)

	//2. call service level to handle business logic

	//3. response
}