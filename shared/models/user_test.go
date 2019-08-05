package models_test

import (
	"testing"

	config "github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

var userService = models.UserDbService{models.Connect(config.LoadConfig())}

func TestGetUserByID(t *testing.T) {

	defer userService.DB.Close()
	user := models.User{
		ID:        "cjyy6zsvl000001ladldsby9z",
		Firstname: "Enaho",
		Lastname:  "Murphy",
		Email:     "enahomurphy@gmail.com",
		Username:  "enahomurphy",
		Password:  "testing",
	}

	userService.DB.Create(&user)

	testCases := []struct {
		Title    string
		ID       string
		Response *models.User
		Error    string
	}{
		{
			Title:    "When a user does not exist",
			ID:       "cjyy7054h000101lacboe5ve9",
			Response: &models.User{},
			Error:    "User with id 44 cannot be found",
		},
		{
			Title:    "When a user exist",
			ID:       "cjyy7054h000101lacboe5vee",
			Response: &models.User{},
			Error:    "",
		},
	}

	for _, testCase := range testCases {
		user, err := userService.GetUserByID(testCase.ID)

		if err != nil {
			if testCase.Error != err.Error() {
				t.Error("expected", testCase.Error)
				t.Error("actual", err.Error())
			}
		} else {
			if user.ID != testCase.ID {
				t.Errorf("expected %d to equal %d", user.ID, testCase.Response.ID)
				t.Errorf("expected %s to equal %s", user.Firstname, testCase.Response.Lastname)
				t.Errorf("expected %s to equal %s", user.Username, testCase.Response.Username)
				t.Errorf("expected %s to equal %s", user.Password, testCase.Response.Password)
			}
		}
	}

	userService.DB.Exec("Truncate table users")
}
