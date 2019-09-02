package resolver

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

func (r *queryResolver) GetAsset(ctx context.Context, id string) (*models.Asset, error) {
	var asset models.Asset

	svc := ctx.Value(config.CTXKeyservices).(*service.Services)
	asset, err := svc.Datastore.AssetDB.GetAsset(id)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (m *mutationResolver) CreatePresignedURL(ctx context.Context, input []*models.AssetInput) ([]*models.AssetUploadOutput, error) {
	var resp []*models.AssetUploadOutput

	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	var assets []models.Asset

	for _, file := range input {
		mediaType := file.MediaType

		if mediaType == "" {
			mediaType = strings.Split(file.MimeType, "/")[1]
		}

		key := getAssetKey(file.ID, file.ID, file.MimeType)
		signedURL, _ := service.AWS.GetS3SignedURL(key, file.MimeType, time.Duration(5*time.Minute))
		assetURL := getAssetURL(signedURL, mediaType)

		asset := models.Asset{
			ID:        file.ID,
			MediaType: file.MediaType,
			MimeType:  file.MimeType,
			Filename:  file.Filename,
			Size:      file.Size,
			Key:       key,
			Url:       assetURL,
			User: []models.User{
				{
					ID: user.ID,
				},
			},
		}

		uploadedAsset := models.AssetUploadOutput{
			ID:        file.ID,
			AssetURL:  assetURL,
			UploadURL: signedURL,
		}

		resp = append(resp, &uploadedAsset)
		assets = append(assets, asset)
	}

	_, err := service.Datastore.AssetDB.CreateAssets(assets)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func getAssetKey(id, assetId, mimeType string) string {
	year, month, _ := time.Now().Date()
	mime := strings.Split(mimeType, "/")
	key := fmt.Sprintf(
		"vendors/%s/%s/%d/%s/%s.%s",
		id,
		mime[0],
		year,
		strings.ToLower(month.String()),
		assetId,
		mime[1],
	)

	return key
}

func getAssetURL(signedURL, mediaType string) string {
	subString := fmt.Sprintf(".%s?", mediaType)
	url := strings.Split(signedURL, subString)[0]
	return fmt.Sprintf("%s.%s", url, mediaType)
}
