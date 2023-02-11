package main

import (
	"fmt"
	"net/http"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	realIp := getIp(r)
	w.Header().Add("Cache-Control", "no-store")
	fmt.Fprintf(w, `ip_addr: %v
user_agent: %v
method: %v
forwarded: %v`,
		realIp, r.Header.Get("user-agent"), r.Method, r.Header.Get("X-Forwarded-For"))
}
