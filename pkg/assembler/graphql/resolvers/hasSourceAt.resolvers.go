package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestHasSourceAt is the resolver for the ingestHasSourceAt field.
func (r *mutationResolver) IngestHasSourceAt(ctx context.Context, pkg model.PkgInputSpec, pkgMatchType model.MatchFlags, source model.SourceInputSpec, hasSourceAt model.HasSourceAtInputSpec) (string, error) {
	ingestedHasSourceAt, err := r.Backend.IngestHasSourceAt(ctx, pkg, pkgMatchType, source, hasSourceAt)
	if err != nil {
		return "", err
	}
	return ingestedHasSourceAt.ID, err
}

// IngestHasSourceAts is the resolver for the ingestHasSourceAts field.
func (r *mutationResolver) IngestHasSourceAts(ctx context.Context, pkgs []*model.PkgInputSpec, pkgMatchType model.MatchFlags, sources []*model.SourceInputSpec, hasSourceAts []*model.HasSourceAtInputSpec) ([]string, error) {
	funcName := "IngestHasSourceAts"
	if len(pkgs) != len(sources) {
		return []string{}, gqlerror.Errorf("%v :: uneven packages and sources for ingestion", funcName)
	}
	if len(pkgs) != len(hasSourceAts) {
		return []string{}, gqlerror.Errorf("%v :: uneven packages and hasSourceAt for ingestion", funcName)
	}
	return r.Backend.IngestHasSourceAts(ctx, pkgs, &pkgMatchType, sources, hasSourceAts)
}

// HasSourceAt is the resolver for the HasSourceAt field.
func (r *queryResolver) HasSourceAt(ctx context.Context, hasSourceAtSpec model.HasSourceAtSpec) ([]*model.HasSourceAt, error) {
	return r.Backend.HasSourceAt(ctx, &hasSourceAtSpec)
}
