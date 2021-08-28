package main

import (
	"fmt"
	"net/http"

	"github.com/pmadhvi/loan-comparator/router"
)

func main() {
	fmt.Println("Hello welcome to loan comparator!")

	http.Handle("/", router.Router())

	// start the http server on port 3000 and pass the chi router as handler
	http.ListenAndServe(":3000", nil)

}

// func ApplicationCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx := context.WithValue(r.Context(), "uuid", uuid.UUID)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
