package middlewares

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Aditya7880900936/postgres-golang/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type response struct {
	StockID int64  `json:"stockid,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")
	return db

}

func CreateStock(w http.ResponseWriter, r *http.Request){
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err!= nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}

	insertID := insertStock(stock)
    
	res := response{
		StockID: insertID,
		Message: "Stock Created Successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	id , err := strconv.Atoi(params["id"])
	if err!= nil {
		log.Fatalf("Unable to convert the string into int %v", err)
	}

	stock, err := getStock(int64(id))
	if err!= nil {
		log.Fatalf("Unable to get stock %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

// func GetAllStock()
// func UpdateStock()
// func DeleteStock()
