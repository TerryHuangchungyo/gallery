package controller

import (
	"encoding/json"
	"gallery-frontend/repository"
	"net/http"
	"strconv"
)

var Mux *http.ServeMux

func init() {
	Mux = http.NewServeMux()

	Mux.Handle("/", http.FileServer(http.Dir("/go/src/assets")))

	Mux.HandleFunc("/api/exhibition", GetExhibition)
	Mux.HandleFunc("/api/exhibition/list", GetExhibitionList)
}

func GetExhibition(writer http.ResponseWriter, request *http.Request) {
	// 預設回傳格式
	result := map[string]interface{}{
		"status":  "ok",
		"message": "",
		"data":    nil,
	}

	defer func() {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		response, _ := json.Marshal(result)
		writer.Write(response)
	}()

	urlValues := request.URL.Query()
	exhibitionID, err := strconv.ParseInt(urlValues.Get("id"), 10, 64)
	if err != nil {
		result["message"] = "Query Parameter Error."
		return
	}

	exhibition, err := repository.GetExhibitionInfoByExhibitionID(exhibitionID)
	if err != nil {
		result["message"] = "Server Internel Error."
		return
	}

	paintList, err := repository.GetExhibitionPaintInfoByExhibitionID(exhibitionID)
	if err != nil {
		result["message"] = "Server Internel Error."
		return
	}

	exhibition.Paints = paintList
	result["data"] = exhibition
	return
}

func GetExhibitionList(writer http.ResponseWriter, request *http.Request) {
	// 預設回傳格式
	result := map[string]interface{}{
		"status":  "ok",
		"message": "",
		"data":    nil,
	}

	defer func() {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		response, _ := json.Marshal(result)
		writer.Write(response)
	}()

	exhibitionList, err := repository.GetExhibitionList()

	if err != nil {
		result["message"] = "Server Internel Error."
		return
	}

	result["data"] = exhibitionList
	return
}
