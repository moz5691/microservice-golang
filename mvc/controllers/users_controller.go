package controllers

import (
	"encoding/json"
	"github.com/moz5691/microservices/mvc/services"
	"github.com/moz5691/microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdParam := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiError := &utils.AppliationErrors{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiError)
		w.WriteHeader(apiError.StatusCode)
		w.Write(jsonValue)
		return
	}
	user, apiError := services.GetUser(uint64(userId))

	if apiError != nil {
		w.WriteHeader(apiError.StatusCode)
		jsonValue, _ := json.Marshal(apiError)
		w.Write(jsonValue)
		return
	}
	// return user to client
	jsonValue, _ := json.MarshalIndent(user, "", "    ")
	w.Write(jsonValue)
}
