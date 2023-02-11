package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	httpListen := fmt.Sprintf(":%v", getEnv("PORT", "8000"))
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/json", jsonHandle)
	log.Println("Server Listening on", httpListen)
	log.Fatal(http.ListenAndServe(httpListen, logRequest(http.DefaultServeMux)))
}

func getEnv(key, fallback string) string {
	if val, isFound := os.LookupEnv(key); isFound {
		return val
	}
	return fallback
}

func errorResponse(w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Write([]byte(string("Invalid Request")))
}

func getIp(r *http.Request) string {
	if r.Header.Get("X-Real-IP") == "" {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		return host
	} else {
		return r.Header.Get("X-Real-IP")
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
