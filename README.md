# go-mongo

Go with mongodb for beginners.

Golang is a popular programming language known for its efficiency and concurrency support. MongoDB is a widely-used NoSQL database. Integrating Go with MongoDB allows you to build efficient and scalable applications. Below, I'll outline the steps to connect Go with MongoDB and perform basic operations.

## Run
```bash
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson

go mod init github.com/sharantechuser/gomongo
go mod tidy
go build

```

## Get all students
####  curl http://localhost:4000/api/students



## Create Student


#### http://localhost:4000/api/student

```bash
{
  "name": "Test", 
  "username": "test", 
  "password": "test@123", 
  "address": {
    "address": "Bangalore", 
    "city": "Bangalore", 
    "pin": "123456"
  }
}
```

## Update Student

#### curl -i -X PUT http://localhost:4000/api/student/651dce8ee500ba78ea40b59d

```bash
{
    "_id": "651dce8ee500ba78ea40b59d",
    "name": "Test",
    "username": "testupdate",
    "password": "test#123",
    "address": {
      "address": "New York city",
      "city": "New York",
      "pin": "123456"
    }
  }
```
## Delete Student

#### curl -i -X DELETE http://localhost:4000/api/student/651dce8ee500ba78ea40b59d


