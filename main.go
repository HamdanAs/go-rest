package main

import (
	"log"
	"net/http"

	"github.com/HamdanAs/goRest/controllers"
	"github.com/HamdanAs/goRest/database"
	"github.com/HamdanAs/goRest/models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()

	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	initializeHandlers(router)

	log.Fatal(http.ListenAndServe(":8090", router))
}

func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletePerson).Methods("DELETE")

}

func initDB() {
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "root",
		DB:         "go_learn",
	}

	connectionString := database.GetConnectionString(config)

	err := database.Connect(connectionString)

	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&models.Person{})
}
