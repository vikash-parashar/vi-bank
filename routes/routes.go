package routes

import (
	"database/sql"
	"go-bank/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Register routes for bank operations
	router.HandleFunc("/create-account", handlers.CreateAccountHandler(db)).Methods("POST")
	router.HandleFunc("/delete-account/{id}", handlers.DeleteAccountHandler(db)).Methods("DELETE")
	router.HandleFunc("/update-account/{id}", handlers.UpdateAccountHandler(db)).Methods("PUT")
	router.HandleFunc("/get-account/{id}", handlers.GetAccountHandler(db)).Methods("GET")
	router.HandleFunc("/transaction-history/{id}", handlers.GetTransactionHistoryHandler(db)).Methods("GET")
	router.HandleFunc("/send-money", handlers.SendMoneyHandler(db)).Methods("POST")
	router.HandleFunc("/check-balance/{id}", handlers.CheckBalanceHandler(db)).Methods("GET")

	return router
}
