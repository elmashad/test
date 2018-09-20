package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"subscriber-app/controllers"
)

func main() {

	bookingHandler := controllers.NewBookingHandler()

	r := mux.NewRouter()
	bookingrouter := r.PathPrefix("/booking").Subrouter()
	bookingrouter.Methods("POST").Path("").HandlerFunc(bookingHandler.AddBookingHandler)
	http.ListenAndServe(":8033", r)
}

