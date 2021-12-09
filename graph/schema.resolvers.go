package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"MaxContinuousSum/graph/generated"
	"MaxContinuousSum/graph/model"
	"MaxContinuousSum/maxSum"
	"context"
)

func (r *mutationResolver) Maxsum(ctx context.Context, list []int) (*model.MaxSum, error) {
	var result model.MaxSum
	result.Sum, result.Positions = maxSum.MaxContinuousSum(list)
	return &result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
