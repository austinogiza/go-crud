package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ToDo struct {
	gorm.Model

	ID   uint
	Task string `json:"task"`
}

var db *gorm.DB
var err error

func Routes() {
	r := mux.NewRouter()

	//api routes
	r.HandleFunc("/todo", allToDos).Methods("Get")
	r.HandleFunc("/todo/{id}", toDoDetails).Methods("GET")
	r.HandleFunc("/create", createToDo).Methods("POST")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

func Migrations() {
	// dialect := os.Getenv("DIALECT")

	//database connection string
	dbURI := "host=localhost user=postgres password=austinforreal dbname=CRUD port=5432 sslmode=disable TimeZone=Africa/Lagos"
	//opening database

	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)

		panic("Can not connect to database")
	} else {
		fmt.Println("Database connected successfully")

	}

	// Migrate the schema
	db.AutoMigrate(&ToDo{})
}
func main() {
	Routes()
	Migrations()

}

func allToDos(w http.ResponseWriter, r *http.Request) {

	var todo []ToDo
	db.Find(&todo)

	json.NewEncoder(w).Encode(todo)

}

func toDoDetails(w http.ResponseWriter, r *http.Request) {
	var todo []ToDo
	params := mux.Vars(r)

	db.First(&todo, params["id"])
	json.NewEncoder(w).Encode(todo)
}

func createToDo(w http.ResponseWriter, r *http.Request) {
	var todo []ToDo

	json.NewDecoder(r.Body).Decode(&todo)
	fmt.Println(r.Body)
	db.Create(&todo)

}
