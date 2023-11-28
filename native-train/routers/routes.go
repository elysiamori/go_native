package routers

import (
	"net/http"

	"github.com/elysiamori/go_native/native-train/handler"
	"github.com/gorilla/mux"
)

// InitRouter initializes router
func InitRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", handler.GetDatas).Methods("GET")
	r.HandleFunc("/api/books/{id}", handler.GetDataByID).Methods("GET")
	r.HandleFunc("/api/books/{uuid}", handler.GetDataByUUID).Methods("GET")
	r.HandleFunc("/api/books", handler.AddData).Methods("POST")
	http.Handle("/", r)
}
