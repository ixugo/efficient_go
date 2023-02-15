package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != "POST" {
		// 	w.WriteHeader(400)
		// 	io.WriteString(w, "请求方式有误")
		// 	return
		// }
		w.WriteHeader(200)
		io.WriteString(w, "OK")
	})
	fmt.Println("server OK")
	http.ListenAndServe(":9922", enableCORS(m))

}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add the "Vary: Origin" header.
		w.Header().Add("Vary", "Origin")
		// Get the value of the request's Origin header.
		origin := r.Header.Get("Origin")
		fmt.Println(origin)
		fmt.Println(r.Method)
		// Only run this if there's an Origin request header present.
		if origin != "" {
			// Loop through the list of trusted origins, checking to see if the request
			// origin exactly matches one of them. If there are no trusted origins, then
			// the loop won't be iterated.
			// for i := range app.config.cors.trustedOrigins {
			// 	if origin == app.config.cors.trustedOrigins[i] {
			// 		// If there is a match, then set a "Access-Control-Allow-Origin"
			// 		// response header with the request origin as the value and break
			// 		// out of the loop.
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Add("Content-Type", "application/json")
			// w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE,OPTIONS")
			w.Header().Add("Access-Control-Max-Age", "60")
			w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With, accept, authorization, content-type")
			// 		break
			// 	}
			// }
		}
		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
