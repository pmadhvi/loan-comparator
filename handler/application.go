package handler

import (
	"fmt"
	"net/http"
)

func GetApplicationById(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside getApplicationById")
	fmt.Fprintf(rw, "Inside getApplicationById")
}

func GetApplications(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetApplications")
	fmt.Fprintf(rw, "Inside GetApplications")
}

func CreateApplication(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside CreateApplication")
	fmt.Fprintf(rw, "Inside CreateApplication")
}
