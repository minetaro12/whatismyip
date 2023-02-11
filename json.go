package main

import (
	"encoding/json"
	"net/http"
)

type Res struct {
	Ip_addr    string `json:"ip_addr"`
	User_agent string `json:"user_agent"`
	Method     string `json:"method"`
	Forwarded  string `json:"forwarded"`
}

func jsonHandle(w http.ResponseWriter, r *http.Request) {
	realIp := getIp(r)
	jsonData, err := json.MarshalIndent(
		Res{
			Ip_addr:    realIp,
			User_agent: r.Header.Get("user-agent"),
			Method:     r.Method,
			Forwarded:  r.Header.Get("X-Forwarded-For"),
		}, "", "  ")
	if err != nil {
		errorResponse(w)
		return
	}
	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.Write(jsonData)
}
