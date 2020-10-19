package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/models"
	"server/services"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/*Error is a struct to send errors as json*/
type Error struct {
	Error string `json:"error"`
}

/*CreateMeetingAPI is our type to register*/
type CreateMeetingAPI int

func (ca CreateMeetingAPI) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	database, err := services.MongoClient.ListDatabaseNames(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(database)

	if req.Method != "POST" {
		err := Error{Error: "Invalid HTTP Method for the path"}
		fmt.Println("hello here")

		json.NewEncoder(res).Encode(err)
		return
	}

	var newMeeting models.Meeting

	if err := json.NewDecoder(req.Body).Decode(&newMeeting); err != nil {
		res.WriteHeader(400)
		err := Error{Error: "Invalid json format"}
		json.NewEncoder(res).Encode(err)
		return
	}

	newMeeting.Timestamp = time.Now().Unix()

	fmt.Println(newMeeting)

	collection := services.MongoClient.Database("meetingsAppointfy").Collection("meetings")

	if result, err := collection.InsertOne(context.TODO(), newMeeting); err != nil {
		res.WriteHeader(400)
		err := Error{Error: "Somthing broke"}
		json.NewEncoder(res).Encode(err)
	} else {
		res.WriteHeader(201)
		json.NewEncoder(res).Encode(result)
	}

	// io.WriteString(res, "Hello world. This acutall works")
}

//Writing code for getting meeting
// var meeting models.Meeting

// 		if err := collection.FindOne(context.TODO(), models.Meeting{ID: stringObjectID}).Decode(&meeting); err != nil {
// 			res.WriteHeader(400)
// 			err := Error{Error: "Somthing broke"}
// 			json.NewEncoder(res).Encode(err)
// 			return
// 		} else {

// 		}
