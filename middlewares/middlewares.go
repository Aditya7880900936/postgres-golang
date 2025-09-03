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

func GetAllStock(w http.ResponseWriter, r *http.Request){
	stocks, err := getAllStock()
	if err!= nil {
		log.Fatalf("Unable to get all stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		log.Fatalf("Unable to convert the string into int %v", err)
	}
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err!= nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}

	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock updated successfully %v", updatedRows)

	res := response{
		StockID: int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		log.Fatalf("Unable to convert the string into int %v", err)
	}

	deletedRows := deleteStock(int64(id))

	msg := fmt.Sprintf("Stock deleted successfully %v", deletedRows)

	res := response{
		StockID: int64(id),
		Message: msg,
	}	

	json.NewEncoder(w).Encode(res)
}
