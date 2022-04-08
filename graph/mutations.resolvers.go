package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"demo/db"
	"demo/graph/generated"
	"demo/graph/model"
	"demo/types"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*types.User, error) {
	user := types.User{Name: input.Name}
	err := db.DB.Create(&user).Error
	return &user, err
}

func (r *mutationResolver) CreateScreech(ctx context.Context, input model.NewScreech) (*types.Screech, error) {
	user := types.User{}
	err := db.DB.Where("id = ?", input.UserID).First(&user).Error
	if err != nil {
		return nil, err
	}

	screech := types.Screech{
		UserID:  user.ID,
		Content: input.Content,
	}

	err = db.DB.Create(&screech).Error
	return &screech, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
