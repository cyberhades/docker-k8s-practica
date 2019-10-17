package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey []byte

type credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func create(w http.ResponseWriter, r *http.Request) {

	var creds credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}

func main() {
	file, err := os.Open("/usr/share/key/key.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jwtKey, err = ioutil.ReadAll(file)
	http.HandleFunc("/create", create)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
