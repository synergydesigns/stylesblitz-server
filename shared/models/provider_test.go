package models_test

import (
	"testing"

	config "gitlab.com/synergy-designs/style-blitz/shared/config"
	"gitlab.com/synergy-designs/style-blitz/shared/models"
)

func TestGetUserByID(t *testing.T) {
	
	// seed providers 
	// seed services
	

	for _, testCase := range testCases {
		user, err := db.GetUserByID(testCase.ID)

		if err != nil {
			if testCase.Error != err.Error() {
				t.Error("expected", testCase.Error)
				t.Error("actual", err.Error())
			}
		} else {
			if user != testCase.Response {
				t.Errorf("expected %d to equal %d", user.ID, testCase.Response.ID)
				t.Errorf("expected %s to equal %s", user.Name, testCase.Response.Name)
				t.Errorf("expected %s to equal %s", user.Username, testCase.Response.Username)
				t.Errorf("expected %s to equal %s", user.Password, testCase.Response.Password)
			}
		}
	}
}
