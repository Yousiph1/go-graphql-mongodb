package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Yousiph1/go-graphql/graph/generated"
	"github.com/Yousiph1/go-graphql/graph/model"
	"github.com/Yousiph1/go-graphql/links"
	"github.com/Yousiph1/go-graphql/users"
	"github.com/Yousiph1/go-graphql/auth"
)


func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	 user := users.User{}
   user.Name = input.Name
	 user.Password = input.Password
	 err := user.Save()
	 token, err := auth.GenerateToken(user.Name)
	 if err != nil {
		 return "", err
	 }
	 return token, nil
}


func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
 user := ctx.Value("user")
 var link links.Link
 link.Address = input.Address
 link.Title = input.Title
 id, err := link.Save(user.(string))
 if err != nil {
	 return nil, err
 }

 return  &model.Link{ID: fmt.Sprint(id),Address: input.Address, Title: input.Title}, nil
}


func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := users.User{}
	user.Name = input.Name
	user.Password = input.Password
	err := user.Login()
  if err != nil {
		return "", err
	}
	token, err := auth.GenerateToken(user.Name)
	if err != nil {
		return "", err
	}
	return token, nil
}


func (r *mutationResolver) GetToken(ctx context.Context, input model.RefreshToken) (string, error) {
	user, err := auth.ParseToken(input.Token)
	if err != nil {
		return "", err
	}
	newToken, err := auth.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return newToken, nil
}


func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
   link := links.Link{}
	 links, err := link.GetAll()
	 if err != nil {
		 return nil , err
	 }
   return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
