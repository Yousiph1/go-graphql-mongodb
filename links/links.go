package links

import (
  "context"
  "fmt"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"

  "github.com/Yousiph1/go-graphql/database"
  "github.com/Yousiph1/go-graphql/graph/model"
  "github.com/Yousiph1/go-graphql/users"
)

var col *mongo.Collection

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func(l *Link) Save(name string) (id interface{}, err error){
  col = database.Client.Database("news").Collection("links")
  userColl := database.Client.Database("news").Collection("users")
  user := model.User{}
  err = userColl.FindOne(context.TODO(), bson.M{"name": name}).Decode(&user)
  res, err  := col.InsertOne(context.Background(),
                   bson.M{"title": l.Title, "address": l.Address,"user":bson.M{"name":name,"id":user.ID}})
  if err == nil {
    id = res.InsertedID
  }
  return
}

func(l *Link) GetAll() ([]*model.Link, error) {
   col = database.Client.Database("news").Collection("links")

   cur, err := col.Find(context.TODO(), bson.D{})

   if err != nil {
     return nil, err
   }
   defer cur.Close(context.Background())

   links := []*model.Link{}
    err =  cur.All(context.TODO(),&links)

     if err != nil {
       return nil, err
     }

    for _, l := range links {
      if l.User != nil {
          fmt.Println(l.User)
        }
    }

   return links, nil
}
