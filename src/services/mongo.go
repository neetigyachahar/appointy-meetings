package services

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoClient has the client*/
var MongoClient *mongo.Client

/*ConnectMongoDB connects to database*/
func ConnectMongoDB() {
	Ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI("mongodb+srv://neetigya:pvx8RHA8CQb8O07l@cluster0-gdzqa.mongodb.net/meetings?retryWrites=true&w=majority"))

	fmt.Println(reflect.TypeOf(Ctx))

	MongoClient = client

	database, err := client.ListDatabaseNames(Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(database)

	if err != nil {
		log.Fatal(err)
	}
}

// func GetClient() *mongo.Client {
// 	return
// }
