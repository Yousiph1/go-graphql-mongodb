package users

import (
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "golang.org/x/crypto/bcrypt"
  "go.mongodb.org/mongo-driver/bson"

  "github.com/Yousiph1/go-graphql/graph/model"
  "github.com/Yousiph1/go-graphql/database"
)
var col *mongo.Collection
type User struct {
	ID       string  `bson:"_id"  json:"id"`
	Name     string  `bson:"name" json:"name"`
  Password string  `bson:"password" json:"password"`
}

func (u *User) Save() (err error) {
  col = database.Client.Database("news").Collection("users")
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),14)
  if err != nil {
    return
  }
   _, err = col.InsertOne(context.Background(),bson.M{"name": u.Name, "password": hashedPassword})
  return
}

func (u *User) Login() (err error) {
   user := model.Login{}
   col = database.Client.Database("news").Collection("users")
   col.FindOne(context.TODO(),bson.M{"name": u.Name}).Decode(&user)

   err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(u.Password))

   return
}
