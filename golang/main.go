package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/username/schedule/helper"
	"github.com/username/schedule/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

func meetingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getMeeting(w, r)
		return
	case "POST":
		createMeeting(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}
func meetingsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getMeetings(w, r)
		return
	case "POST":
		createMeeting(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func getMeetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Meeting array
	var meetings []models.Meeting

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var meeting models.Meeting
		// & character returns the memory address of the following variable.
		err := cur.Decode(&meeting) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		meetings = append(meetings, meeting)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(meetings) // encode similar to serialize process.
}

func getMeeting(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var meeting models.Meeting

	query := r.URL.Query()
	idParams := query.Get("id")

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(idParams)

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&meeting)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(meeting)
}

func createMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var meeting models.Meeting

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&meeting)

	// insert our meeting model.
	result, err := collection.InsertOne(context.TODO(), meeting)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {

	http.HandleFunc("/meeting", meetingHandler)
	http.HandleFunc("/meetings/", meetingsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
