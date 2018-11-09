package models_test

import (
	"testing"

	config "gitlab.com/synergy-designs/style-blitz/shared/config"
	"gitlab.com/synergy-designs/style-blitz/shared/models"
	"gitlab.com/synergy-designs/style-blitz/shared/seeder"
)

var providerService = models.ProviderDbService{models.Connect(config.LoadConfig())}

func TestGetProvidersByLocationAndService(t *testing.T) {

	seed := seeder.Seeder{}

	seed.Init().LoadData("provider").Seed("provider")
	seed.LoadData("address").Seed("address")
	seed.LoadData("category").Seed("category")
	seed.LoadData("service").Seed("service")

	defer seed.Close()

	testCases := []struct {
		Title             string
		ServiceName       string
		Lat, Long, Radius float64
		Count             int
	}{
		{
			Title: `should  get provider that 
				offers services within 1 km readius`,
			ServiceName: "Extensions",
			Lat:         6.57747305,
			Long:        3.36740283,
			Radius:      1,
			Count:       2,
		},
		{
			Title: `should  get provider that 
				offers services within 1 km radius of ajao estate`,
			ServiceName: "Extensions",
			Lat:         6.54565167,
			Long:        3.33274265,
			Radius:      1,
			Count:       1,
		},
	}

	for _, testCase := range testCases {
		providers, _ := providerService.GetProvidersByServiceAndLocation(
			testCase.ServiceName,
			testCase.Lat,
			testCase.Long,
			testCase.Radius,
		)

		if len(providers) != testCase.Count {
			t.Errorf("expected result count %d to equal %d", len(providers), testCase.Count)
		}
	}

	seed.Clean()
}
