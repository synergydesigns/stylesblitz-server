package resolver

import (
	"context"
	"encoding/json"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type VendorResolver struct{ *Resolver }

func (r *VendorResolver) Addresses(ctx context.Context, obj *models.Vendor) ([]*models.Address, error) {
	panic("not implemented")
}

func (r *VendorResolver) Services(ctx context.Context, obj *models.Vendor) ([]*models.Service, error) {
	panic("not implemented")
}

func (r *VendorResolver) Phone(ctx context.Context, obj *models.Vendor) (*string, error) {
	phoneBytes, err := json.Marshal(obj.Phone)

	phone := string(phoneBytes)

	if err != nil {
		return &phone, err
	}

	return &phone, nil
}
