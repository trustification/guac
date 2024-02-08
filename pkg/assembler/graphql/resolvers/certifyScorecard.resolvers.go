package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestScorecard is the resolver for the ingestScorecard field.
func (r *mutationResolver) IngestScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (string, error) {
	ingestedScorecard, err := r.Backend.IngestScorecard(ctx, source, scorecard)
	if err != nil {
		return "", err
	}
	return ingestedScorecard.ID, err
}

// IngestScorecards is the resolver for the ingestScorecards field.
func (r *mutationResolver) IngestScorecards(ctx context.Context, sources []*model.SourceInputSpec, scorecards []*model.ScorecardInputSpec) ([]string, error) {
	funcName := "IngestScorecards"
	ingestedScorecardsIDS := []string{}
	if len(sources) != len(scorecards) {
		return ingestedScorecardsIDS, fmt.Errorf("%v :: uneven source and scorecards for ingestion", funcName)
	}
	ingestedScorecards, err := r.Backend.IngestScorecards(ctx, sources, scorecards)
	if err == nil {
		for _, scorecard := range ingestedScorecards {
			ingestedScorecardsIDS = append(ingestedScorecardsIDS, scorecard.ID)
		}
	}
	return ingestedScorecardsIDS, err
}

// Scorecards is the resolver for the scorecards field.
func (r *queryResolver) Scorecards(ctx context.Context, scorecardSpec model.CertifyScorecardSpec) ([]*model.CertifyScorecard, error) {
	return r.Backend.Scorecards(ctx, &scorecardSpec)
}
