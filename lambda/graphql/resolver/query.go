package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type queryResolver struct{ *Resolver }

func (resolver *queryResolver) GetSuggestions(ctx context.Context, query string) ([]*models.Autocomplete, error) {
	panic(0)
}
