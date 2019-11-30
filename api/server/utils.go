package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondSuccess(w http.ResponseWriter, status int, v interface{}) {
	w.WriteHeader(status)
	if v == nil {
		return
	}
	body, err := json.Marshal(v)
	if err != nil {
		body, ok := v.([]byte)
		if ok {
			w.Write(body)
			return
		}
		strBody, ok := v.(string)
		if ok {
			w.Write([]byte(strBody))
			return
		}
		return
	}
	w.Write(body)
}

func respondError(w http.ResponseWriter, err error, status int, v interface{}) {
	w.WriteHeader(status)
	log.Println(err)
	if v == nil {
		return
	}
	body, err := json.Marshal(v)
	if err != nil {
		body, ok := v.([]byte)
		if ok {
			w.Write(body)
			return
		}
		strBody, ok := v.(string)
		if ok {
			w.Write([]byte(strBody))
			return
		}
		return
	}
	w.Write(body)
}
