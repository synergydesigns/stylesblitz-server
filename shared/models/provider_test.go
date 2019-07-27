package models_test

// import (
// 	"testing"

// 	config "github.com/synergydesigns/stylesblitz-server/shared/config"
// 	"github.com/synergydesigns/stylesblitz-server/shared/models"
// 	"github.com/synergydesigns/stylesblitz-server/shared/seeder"
// )

// var VendorService = models.VendorDbService{models.Connect(config.LoadConfig())}

// func TestGetVendorsByLocationAndService(t *testing.T) {

// 	seed := seeder.Seeder{}

// 	seed.Init().LoadData("Vendor").Seed("Vendors")
// 	seed.LoadData("address").Seed("addresss")
// 	seed.LoadData("category").Seed("categories")
// 	seed.LoadData("service").Seed("services")

// 	defer seed.Close()

// 	testCases := []struct {
// 		Title             string
// 		ServiceName       string
// 		Lat, Long, Radius float64
// 		Count             int
// 	}{
// 		{
// 			Title: `should  get Vendor that
// 				offers services within 1 km readius`,
// 			ServiceName: "Extensions",
// 			Lat:         6.57747305,
// 			Long:        3.36740283,
// 			Radius:      1,
// 			Count:       2,
// 		},
// 		{
// 			Title: `should  get Vendor that
// 				offers services within 1 km radius of ajao estate`,
// 			ServiceName: "Extensions",
// 			Lat:         6.54565167,
// 			Long:        3.33274265,
// 			Radius:      1,
// 			Count:       1,
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		Vendors, _ := VendorService.GetVendorsByServiceAndLocation(
// 			testCase.ServiceName,
// 			testCase.Lat,
// 			testCase.Long,
// 			testCase.Radius,
// 		)

// 		if len(Vendors) != testCase.Count {
// 			t.Errorf("expected result count %d to equal %d", len(Vendors), testCase.Count)
// 		}
// 	}

// 	seed.Clean()
// }
