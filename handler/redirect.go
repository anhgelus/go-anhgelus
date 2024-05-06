package handler

import (
	"github.com/anhgelus/go-anhgelus/data"
	"github.com/gorilla/mux"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	slug, ok := v["slug"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	l := data.Cfg.GetLink(slug)
	if len(l) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(w, r, l, http.StatusFound)
}
