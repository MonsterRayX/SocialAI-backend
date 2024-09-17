package main

import (
	"fmt"
	"log"
	"net/http"

	"around/backend"
	"around/handler"
	"around/util"
)
func main() {
    fmt.Println("started-service")

	config, err := util.LoadApplicationConfig("conf", "deploy.yml")
	if err != nil {
		panic(err)
	}
	backend.InitElasticsearchBackend()
	backend.InitGCSBackend()

    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}