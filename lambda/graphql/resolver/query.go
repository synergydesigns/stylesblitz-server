package resolver

import (
	"context"
	"log"
	"regexp"
	"strings"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type queryResolver struct{ *Resolver }

func (resolver *queryResolver) GetSuggestions(ctx context.Context, query string) ([]*models.Autocomplete, error) {
	service := config.GetServices(ctx)

	suggestions, err := service.Datastore.AutocompleteDB.GetSuggestions(query)

	for _, suggestion := range suggestions {
		var reg = regexp.MustCompile(`\s+`)

		log.Println(suggestion.Type)

		if suggestion.Type == "vendors" || suggestion.Type == "categories" {
			cleanQuery := reg.ReplaceAllString(strings.TrimSpace(suggestion.Query), "-")
			suggestion.Url = "/" + suggestion.Type + "/" + suggestion.ID + "-" + cleanQuery
		} else {
			cleanQuery := reg.ReplaceAllString(strings.TrimSpace(suggestion.Query), "-")
			suggestion.Url = "/" + suggestion.Type + "?query=" + cleanQuery
		}
	}

	return suggestions, err
}
