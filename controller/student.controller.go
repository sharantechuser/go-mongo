package controller

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharantechuser/gomongo/model"
	"github.com/sharantechuser/gomongo/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentController struct {
	StudentService service.IStudentService
}

func New(studentservice service.IStudentService) StudentController {

	return StudentController{
		StudentService: studentservice,
	}

}

const (
	connectionString = "localhost:27917"
	db_name          = "gomongo"
	coll_name        = "student"
)

var Student_collection *mongo.Collection

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect the client when done
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Access the "test" database and the "people" collection
	Student_collection = client.Database("gomongo").Collection("student")
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {

	students, err := service.NewStudentService().GetAllStudents()
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	bytes, err := json.Marshal(students)
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusCreated)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	payload, _ := io.ReadAll(r.Body)

	var student model.Student
	json.Unmarshal(payload, &student)

	err := service.NewStudentService().CreateStudent(student)
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)

}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student

	params := mux.Vars(r)
	_id := params["id"]

	payload, _ := io.ReadAll(r.Body)
	json.Unmarshal(payload, &student)

	_, err := service.NewStudentService().UpdateStudent(_id, student)
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id := params["id"]

	_, err := service.NewStudentService().DeleteStudent(_id, model.Student{})
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
