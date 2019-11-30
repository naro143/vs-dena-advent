package handler

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/log"
)

func respondSuccess(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
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

func respondError(w http.ResponseWriter, r *http.Request, err error, status int, v interface{}) {
	w.WriteHeader(status)
	log.Errorf(r.Context(), err.Error())
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
