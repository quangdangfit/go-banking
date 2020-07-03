package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-banking/helpers"
	"go-banking/interfaces"
	"go-banking/transactions"
	"go-banking/useraccounts"
	"go-banking/users"
	"log"
	"net/http"
)

func transaction(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	auth := r.Header.Get("Authorization")
	var formattedBody interfaces.TransactionBody
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	transaction := useraccounts.Transaction(formattedBody.UserId, formattedBody.From, formattedBody.To, formattedBody.Amount, auth)
	helpers.Response(transaction, w)
}

func getMyTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userID"]
	auth := r.Header.Get("Authorization")

	transactions := transactions.GetMyTransactions(userId, auth)
	helpers.Response(transactions, w)
}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", users.Login).Methods("POST")
	router.HandleFunc("/register", users.Register).Methods("POST")
	router.HandleFunc("/transaction", transaction).Methods("POST")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
