package main

import (
	"github.com/elmashad/test/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	bookingHandler := controllers.NewBookingHandler()

	r := mux.NewRouter()
	bookingrouter := r.PathPrefix("/booking").Subrouter()
	bookingrouter.Methods("POST").Path("").HandlerFunc(bookingHandler.AddBookingHandler)
	http.ListenAndServe(":8033", r)
}
