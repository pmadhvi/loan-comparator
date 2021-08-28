package handler

import (
	"fmt"
	"net/http"
)

func GetJobByApplicationId(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside getJobByApplicationId")
	fmt.Fprintf(rw, "Inside getJobByApplicationId")
}
