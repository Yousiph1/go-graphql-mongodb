package database

import (
    "os"
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Init(){
  //connect to mongodb server
  client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("[Connection to mongo db failed]: make sure 'MONGOURI' is added to .env file")
	}

 //To disconnect from database when the server fails or is terimnated
  Client = client
}
