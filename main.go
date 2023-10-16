package main

import (
	"log"
	"net/http"
	"vibank/storage"

	"github.com/gorilla/mux"
)

func init() {
	if err := storage.CreateTablesFromFile("schema.sql", "user=postgres password=postgres dbname=V!Bank sslmode=disable"); err != nil {
		log.Fatalf("Error: %v\n", err)
		return
	}
}

type Response struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck)
	r.HandleFunc("/user/account/get", GetAccountByUser)
	r.HandleFunc("/user/account/get-all", GetAllAccount) //TODO: auth needed only admin can get all accounts
	r.HandleFunc("/user/account/create", CreateAccount)
	r.HandleFunc("/user/account/update", UpdateAccount)
	r.HandleFunc("/user/account/delete", DeleteAccount) //TODO: auth needed only admin can get all accounts

	http.ListenAndServe(":8080", r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("application is running . . . "))
}

func GetAccountByUser(w http.ResponseWriter, r *http.Request) {

}

func GetAllAccount(w http.ResponseWriter, r *http.Request) {

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {

}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

}
