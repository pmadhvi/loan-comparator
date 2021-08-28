package handler

import (
	"fmt"
	"net/http"
)

func GetApplicationById(rw http.ResponseWriter, r *http.Request) error {
	fmt.Println("Inside getApplicationById")
	return nil
}
