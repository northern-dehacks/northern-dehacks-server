package controllers

import (
	"context"
	"net/http"

	"github.com/northern-dehacks/northern-dehacks-server/database"
)

type signup struct {
	Email string `bson:"email,omitempty"`
}

func SignUpForNewsletter(w http.ResponseWriter, r *http.Request) {
	var signUpDoc signup

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	signUpDoc.Email = r.FormValue("email")

	if err := database.Database.Insert("newsletter", context.Background(), signUpDoc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
