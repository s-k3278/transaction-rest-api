package main

import (
	"net/http"
	//"github.com/gin-gonic/gin"
	"encoding/json"
	"log"

	//	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type employee struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceNAME := "root:Sujit@8723@tcp(localhost:3306)/"
	db, err = gorm.Open("mysql", dataSourceNAME)

	if err != nil {
		panic(err)
	}

	//	db.Exec("CREATE DATABASE employeebalance_db")
	db.Exec("USE employeebalance_db")
	db.AutoMigrate(&employee{})
}

func postEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee employee
	json.NewDecoder(r.Body).Decode(&newEmployee)
	db.Create(&newEmployee)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmployee)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	var UpdatedEmployee employee
	json.NewDecoder(r.Body).Decode(&UpdatedEmployee)
	db.Save(&UpdatedEmployee)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UpdatedEmployee)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create_employees", postEmployee).Methods("POST")
	router.HandleFunc("/update_employees", updateEmployee).Methods("PUT")
	initDB()
	log.Fatal(http.ListenAndServe(":6000", router))
}
