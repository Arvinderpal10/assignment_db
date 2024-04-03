package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// the message format used
type Data struct {
	Message string `json:"message"`
}

// DB Configuration
const (
	connectionString = "mongodb://localhost:27017"
	dbName           = "astra"
	collectionName   = "assignment"
)

// Post Request Function
func postData(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	go insertDataToDB(data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received successfully"))
}

func insertDataToDB(data Data) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database(dbName).Collection(collectionName)
	document := bson.M{
		"message":   data.Message,
		"timestamp": time.Now(),
	}
	_, err = collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/post", postData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
