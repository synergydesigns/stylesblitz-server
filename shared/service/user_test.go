package service_test

import (
	"fmt"
	"testing"

	"gitlab.com/synergy-designs/style-blitz/shared/db"
	"gitlab.com/synergy-designs/style-blitz/shared/service"
)

func TestGetUserByID(t *testing.T) {
	user := db.User{
		Name:     "Enaho Murphy",
		Email:    "enahomurphy@gmail.com",
		Username: "enahomurphy",
		Password: "testing",
	}
	service.DB.Create(&user)

	testCases := []struct {
		Title    string
		ID       uint
		Response db.User
		Error    string
	}{
		{
			Title:    "When a user does not exist",
			ID:       44,
			Response: db.User{},
			Error:    "User with id 44 cannot be foundd",
		},
	}

	for _, testCase := range testCases {

		user, err := service.GetUserByID(testCase.ID)

		if err != nil {
			if testCase.Error != err.Error() {
				t.Error("expected", testCase.Error, "=====but got====", err.Error())
			}
		} else {
			if user != testCase.Response {
				fmt.Println("sccxcxcx")
			}
		}

	}

}
