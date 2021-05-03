package controller

import (
	"encoding/json"
	"gallery-backend/repository"
	. "gallery-backend/utils"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var Mux *http.ServeMux

func init() {
	Mux = http.NewServeMux()

	Mux.Handle("/", http.FileServer(http.Dir("/go/src/assets")))
	Mux.HandleFunc("/api/exhibition", CreateExhibition)
	Mux.HandleFunc("/api/exhibition/list", GetExhibitionList)
	Mux.HandleFunc("/api/exhibition/paint", AddPaintToExhibition)
	Mux.HandleFunc("/api/exhibition/paint/list", GetPaintListByExhibitionID)
	Mux.HandleFunc("/api/paint", AddPaint)
	Mux.HandleFunc("/api/paint/list", GetPaintList)
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

	exhibitionList, _ := repository.GetExhibitionList()

	for index := range exhibitionList {
		paintList, _ := repository.GetExhibitionPaintInfoByExhibitionID(exhibitionList[index].ID)
		exhibitionList[index].Paints = paintList
	}

	result["data"] = exhibitionList
	return
}

func GetPaintListByExhibitionID(writer http.ResponseWriter, request *http.Request) {
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
	exhibitionID, err := strconv.ParseInt(urlValues.Get("exhibition_id"), 10, 64)
	if err != nil {
		result["message"] = "Query Parameter Error."
		return
	}

	paintList, err := repository.GetExhibitionPaintInfoByExhibitionID(exhibitionID)
	if err != nil {
		result["message"] = "Server Internel Error."
		return
	}

	result["data"] = paintList
	return
}

func CreateExhibition(writer http.ResponseWriter, request *http.Request) {
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

	requestPayload, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result["status"] = "no"
		result["message"] = "Error on read request body"
		return
	}

	var requestParam map[string]string
	err = json.Unmarshal(requestPayload, &requestParam)
	if err != nil {
		ErrorLog.Println(err)
		result["status"] = "no"
		result["message"] = "Error on parse request body"
		return
	}

	title, exist := requestParam["title"]
	description, exist := requestParam["description"]
	if !exist {
		result["status"] = "no"
		result["message"] = "Lack of request parameter"
		return
	}

	lastInsertID, err := repository.CreateNewExhibition(title, description)
	if err != nil {
		result["status"] = "no"
		result["message"] = "Internel Server error"
		return
	}

	responsePayload := map[string]interface{}{
		"id":          lastInsertID,
		"title":       title,
		"description": description,
	}
	result["data"] = responsePayload

	return
}

func AddPaintToExhibition(writer http.ResponseWriter, request *http.Request) {
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

	requestPayload, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result["status"] = "no"
		result["message"] = "Error on read request body"
		return
	}

	requestParam := struct {
		ExhibitionID int64   `json:"exhibition_id"`
		PaintIDList  []int64 `json:"paint_id_list"`
	}{}

	err = json.Unmarshal(requestPayload, &requestParam)
	if err != nil {
		ErrorLog.Println(err)
		result["status"] = "no"
		result["message"] = "Error on parse request body"
		return
	}

	err = repository.DeletePaintFromExhibition(requestParam.ExhibitionID)
	if err != nil {
		result["status"] = "no"
		result["message"] = "Internel Server error"
		return
	}

	err = repository.AddPaintToExhibition(requestParam.ExhibitionID, requestParam.PaintIDList)
	if err != nil {
		result["status"] = "no"
		result["message"] = "Internel Server error"
		return
	}

	return
}

func GetPaintList(writer http.ResponseWriter, request *http.Request) {
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

	paintList, _ := repository.GetPaintList()

	result["data"] = paintList
	return
}

func AddPaint(writer http.ResponseWriter, request *http.Request) {
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

	// limit multipartform total 20MB
	err := request.ParseMultipartForm(20 * 1 << 20)
	if err != nil {
		ErrorLog.Printf("PostPaint Error: %v", err)
		result["status"] = "no"
		result["message"] = "Upload image error."
		return
	}

	multiPartForm := request.MultipartForm
	paintName := multiPartForm.Value["name"][0]
	paintFile := multiPartForm.File["image"][0]

	lastInsertId, err := repository.CreateNewPaint(paintName)
	if err != nil {
		result["status"] = "no"
		result["message"] = "Upload image error."
		return
	}

	storedPath := "/go/src/assets/image/"
	imageFileName := strconv.FormatInt(lastInsertId, 10)
	imageFileExtension := filepath.Ext(paintFile.Filename)
	imageFullName := imageFileName + imageFileExtension

	imageFile, err := os.OpenFile(storedPath+imageFullName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		ErrorLog.Printf("Store image file error: %v", imageFile)
		result["status"] = "no"
		result["message"] = "Upload image error."
		return
	}

	uploadImageFile, err := paintFile.Open()
	if err != nil {
		ErrorLog.Printf("Store image file error")
		result["status"] = "no"
		result["message"] = "Upload image error."
		return
	}

	imageData, err := ioutil.ReadAll(uploadImageFile)
	if err != nil {
		ErrorLog.Printf("Store image file error")
		result["status"] = "no"
		result["message"] = "Upload image error."
		return
	}

	_, err = imageFile.Write(imageData)
	if err != nil {
		ErrorLog.Printf("Store image file error: %v", err)
		result["status"] = "no"
		result["message"] = "Upload image error."
		return
	}
	imageFile.Close()

	err = repository.UpdatePaintImageUrl(lastInsertId, imageFullName)
	if err != nil {
		ErrorLog.Printf("Store image file error: %v", err)
		result["status"] = "no"
		result["message"] = "Upload image error."
		return
	}

	return
}
