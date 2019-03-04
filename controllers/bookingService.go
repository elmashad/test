package controllers

import (
	"fmt"
	"net/http"
)

type bookingService struct{}

func NewBookingHandler() *bookingService {
	return &bookingService{}
}

func (eh *bookingService) AddBookingHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("test commit")
	//fmt.Fprint(w, `{"id":%d}`, 2222)
}
