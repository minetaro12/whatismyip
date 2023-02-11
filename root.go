package main

import (
	"fmt"
	"net/http"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	realIp := getIp(r)
	fmt.Fprintf(w, "ip_addr: %v\nuser_agent: %v\nmethod: %v", realIp, r.Header.Get("user-agent"), r.Method)
}
