package controllers

import (
	"encoding/json"
	"golang-api/models"
	u "golang-api/utils"
	"net/http"
)

var Test = func(w http.ResponseWriter, r *http.Request) {

	test := &models.Test{}
	err := json.NewDecoder(r.Body).Decode(test)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := test.Test()
	u.Respond(w, resp)
}
