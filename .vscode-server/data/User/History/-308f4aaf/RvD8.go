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

	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode user data from client %v\n", err)
		return
	}

	success, err := service.CheckUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to read user from Elasticsearch", http.StatusInternalServerError)
		fmt.Printf("Failed to read user from Elasticsearch %v\n", err)
		return 
	}

	if !success{
		http.Error(w, "User doesn't exists or wrong password", http.StatusUnauthorized)
		fmt.Printf("User doesn't exists or wrong password\n")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp": time.Now().Add(time.Hour*24).Unix(),
	})

	toeknString, err := token.SignedString(mySigningKey)
	if err != nil {
		http.
	}
}