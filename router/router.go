package router

import (
	"github.com/gorilla/mux"
	"github.com/sharantechuser/gomongo/controller"
	"net/http"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.Handle("/api/students", http.HandlerFunc(controller.GetAllStudents)).Methods("GET")
	router.Handle("/api/student", http.HandlerFunc(controller.CreateStudent)).Methods("POST")
	router.Handle("/api/student/{id}", http.HandlerFunc(controller.UpdateStudent)).Methods("PUT")
	router.Handle("/api/student/{id}", http.HandlerFunc(controller.DeleteStudent)).Methods("DELETE")

	return router
}
