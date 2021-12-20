package main

import (
	"log"
	"os"
	"context"
  "github.com/Yousiph1/go-graphql/auth"
  "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Yousiph1/go-graphql/graph"
	"github.com/Yousiph1/go-graphql/graph/generated"
  "github.com/Yousiph1/go-graphql/database"
)


//Since I'm using the gin framework, the route controler must be of type gin.HandlerFunc
//This is why we wrap the to return gin.HanlderFunc and extracting the response writer and request from gin.Contenxt
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
  return func (c *gin.Context)  {
  	h.ServeHTTP(c.Writer, c.Request)
  }
}

//playground handler
func playgroundHandler() gin.HandlerFunc {
	 h := playground.Handler("Graphql", "/query")
	 return func (c *gin.Context)  {
	 	  h.ServeHTTP(c.Writer, c.Request)
	 }
}

func main() {
	//load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}
	//no default port to enforce adding port to the .env file
	port := os.Getenv("PORT")

  database.Init()
	defer func ()  {
			if err = database.Client.Disconnect(context.TODO()); err != nil {
				log.Fatal("Failed to disconnect to db")
			}
	}()
 r := gin.Default()
 r.Use(auth.Authenticate())
 r.GET("/", playgroundHandler())
 r.POST("/query",graphqlHandler())
 r.Run(port)
}
