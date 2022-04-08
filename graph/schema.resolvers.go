package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"demo/db"
	"demo/graph/generated"
	"demo/types"
)

func (r *queryResolver) Users(ctx context.Context) ([]types.User, error) {
	users := make([]types.User, 0)
	err := db.DB.Preload("Screeches").Find(&users).Error
	return users, err
}

func (r *queryResolver) Screeches(ctx context.Context) ([]types.Screech, error) {
	screeches := make([]types.Screech, 0)
	err := db.DB.Find(&screeches).Error
	return screeches, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
