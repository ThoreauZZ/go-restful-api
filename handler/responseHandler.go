package handler

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/ThoreauZZ/go-restful-api/model"
)

func RstHandler(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	rst := model.ResponseModel{Message: "ok", Data: data}
	bytes, err := json.Marshal(rst)
	if err != nil {
		log.Println(err)
	}
	w.Write(bytes)
}
