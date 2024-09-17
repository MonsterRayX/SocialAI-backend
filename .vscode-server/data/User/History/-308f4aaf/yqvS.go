package handler

import(
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
	
	"around/model"
	"around/service"

	jwt "github.com/form3tech-oss/jwt-go"
)

var mySigningKey = []byte("secret")

func signingHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received one signing request")
	w.Header().Set("Content-Type", "text/plain")
}