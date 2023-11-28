package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/elysiamori/go_native/native-train/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// 3. Create reads json file
func ReadJSON() (models.Library, error) {

	// 3.1 Open jsonFile
	filePath := filepath.Join("data", "data.json")
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		return models.Library{}, err
	}

	// 3.2  variable to store jsonFile
	var bookData models.Library
	err = json.Unmarshal(jsonFile, &bookData)
	if err != nil {
		return models.Library{}, err
	}

	return bookData, nil
}

// 4 . GetDatas gets all datas
func GetDatas(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Read json file
	jsonData, err := ReadJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set response
	json.NewEncoder(w).Encode(jsonData)
}

// 4.1 GetData gets data by id
func GetDataByID(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// Get params
	params := mux.Vars(r)

	// Read json file
	jsonData, err := ReadJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// Get id from request
	for _, value := range models.Library(jsonData).Library {
		if value.ID == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}

	// Set response
	json.NewEncoder(w).Encode(nil)
}

// 4.2 GetData gets data by uuid
func GetDataByUUID(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// Get params
	params := mux.Vars(r)

	// Read json file
	jsonData, err := ReadJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// Get uuid from parameter url
	for _, value := range models.Library(jsonData).Library {
		if value.Uuid == params["uuid"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}

	// Set response
	json.NewEncoder(w).Encode(nil)
}

// 4.3 AddData adds new data
func AddData(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// Read json file
	jsonData, err := ReadJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// Get request body
	var newBook models.Book
	err = json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the ID
	newBook.Uuid = uuid.New().String()
	newBook.Date = time.Now().Format("2006-01-02")

	// Add new data
	jsonData.Library = append(jsonData.Library, newBook)

	// Write updated data back to the JSON file
	err = writeJSON(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response
	json.NewEncoder(w).Encode(newBook)
}

// writeJSON writes data to the JSON file
func writeJSON(data models.Library) error {
	// Open jsonFile
	filePath := filepath.Join("data", "data.json")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode and write data to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
