package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/mitzukodavis/apicodi/connect"
)
type User struct {
	Username string `json:"username"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
}

func main() {
	connect.InitializeDatabase()
	defer connect.CloseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	log.Println("El servidor en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	//vars := mux.Vars(r)
	//user_id := vars['id']
	user := User{"Eduardo Ismael","test1","test2"}
	json.NewEncoder(w).Encode(user)
}
