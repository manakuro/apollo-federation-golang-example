package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"service-profile/graph/data"
	"service-profile/graph/generated"
	"service-profile/graph/model"
)

func (r *userResolver) Profile(ctx context.Context, obj *model.User) (*model.Profile, error) {
	var res *model.Profile

	for _, p := range data.Profiles {
		if p.UserID == obj.ID {
			res = p
		}
	}
	return res, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
