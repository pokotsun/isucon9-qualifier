package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getSettings(w http.ResponseWriter, r *http.Request) {
	csrfToken := getCSRFToken(r)

	user, _, errMsg := getUser(r)

	ress := resSetting{}
	ress.CSRFToken = csrfToken
	if errMsg == "" {
		ress.User = &user
	}

	ress.PaymentServiceURL = getPaymentServiceURL()

	categories := []Category{}

	err := dbx.Select(&categories, "SELECT * FROM `categories`")
	if err != nil {
		log.Print(err)
		outputErrorMsg(w, http.StatusInternalServerError, "db error")
		return
	}
	ress.Categories = categories

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(ress)
}
