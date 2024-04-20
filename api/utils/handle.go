package utils

import (
	"encoding/json"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) Handler {
	return Handler{
		w: w,
		r: r,
	}
}

type Handler struct {
	w http.ResponseWriter
	r *http.Request
}

func (h Handler) Respond(object any) {
	bytes, err := json.Marshal(object)
	if err != nil {
		return
	}
	_, err = h.w.Write(bytes)
	if err != nil {
		return
	}
}

func (h Handler) RespondError(err error) {
	bytes, err := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	if err != nil {
		return
	}
	_, err = h.w.Write(bytes)
	if err != nil {
		return
	}
}
