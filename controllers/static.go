package controllers

import (
	"net/http"

	"github.com/joncalhoun/lenslocked/views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  tpl.Execute(w, nil)
	}
  }