package resolver

import (
	"context"
	"encoding/json"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type addressResolver struct{ *Resolver }

func (a *addressResolver) Country(ctx context.Context, obj *models.Address) (*string, error) {
	c, err := json.Marshal(obj.Country)

	cString := string(c)

	if err != nil {
		return &cString, err
	}

	return &cString, nil
}

func (a *addressResolver) State(ctx context.Context, obj *models.Address) (*string, error) {
	s, err := json.Marshal(obj.State)

	sString := string(s)

	if err != nil {
		return &sString, err
	}

	return &sString, nil
}
