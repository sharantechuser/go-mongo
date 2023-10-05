package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sharantechuser/gomongo/db"
	"github.com/sharantechuser/gomongo/router"
)

func main() {

	if !db.Initialize() {
		log.Fatal("FATAL! could not connect to database")
	}

	fmt.Println("Go with Mongo API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 4000 ...")
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}

	//--------------
	// // Set client options
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// // Connect to MongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Disconnect the client when done
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// // Access the "test" database and the "people" collection
	// collection := client.Database("gomongo").Collection("student")
	// //--------------------
	// // Insert a document
	// address := model.Address{"New York city", "New York", "123456"}
	// user := model.Student{Name: "John Doe", Username: "test", Password: "test", Address: address}
	// insertResult, err := collection.InsertOne(context.TODO(), user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted document with ID:", insertResult.InsertedID)

	// // Find a document
	// var result model.Student
	// err = collection.FindOne(context.TODO(), bson.M{"name": "John Doe"}).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Found document:", result)
}
