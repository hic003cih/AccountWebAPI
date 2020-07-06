package models

import (
	u "golang-api/utils"
	"os"
	"time"
)

type Test struct {
	Time string `json:"time"`
}

func (test *Test) Test() map[string]interface{} {
	now := time.Now()
	if test.Token == "" {
		test.Time = ""
	}

	if test.Token != "" {
		test.Time = now.String()
	}

	response := u.Message(true, "test")
	response["test"] = test
	return response
}
