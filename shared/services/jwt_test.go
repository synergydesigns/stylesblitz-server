package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	config "github.com/synergydesigns/stylesblitz-server/shared/config"
	models "github.com/synergydesigns/stylesblitz-server/shared/models"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

var JWT *service.JWTService = service.NewJWT(config.LoadConfig())

func TestGenerateAuthToken(t *testing.T) {
	testCase := []struct {
		Title string
		User  models.User
	}{
		{
			Title: "should return a token when GenerateAuthToken is called",
			User:  models.User{},
		},
	}

	for _, test := range testCase {
		if _, err := JWT.GenerateAuthToken(test.User); err != nil {
			t.Errorf("GenerateAuthToken failed expected a valid token but got this %v", err)
		}
	}
}

func TestDecodeToken(t *testing.T) {
	testCases := []struct {
		Title string
		Token string
		Error bool
	}{
		{
			Title: "should user struct when Decode token is called",
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJsb29rYmxpdHoiLCJhdWQiOiJsb2NhbGhvc3Q6MzAwMSIsImV4cCI6MTU5ODM3MDg4OCwibmJmIjoxNTY3MjY4Njg4LCJpYXQiOjE1NjcyNjY4ODgsIlVzZXIiOnsiSUQiOiJjank3bjh3MTEwMDAxbGhxZ2gzNmY2ZHRsIiwiRmlyc3RuYW1lIjoiQWxldGEiLCJMYXN0bmFtZSI6IkNlY2VyZSIsIlVzZXJuYW1lIjoiYWNlY2VyZTAiLCJFbWFpbCI6ImFjZWNlcmUwQHJlZmVyZW5jZS5jb20iLCJQYXNzd29yZCI6IiQyYSQxNCQzZXZHbDdWUWtyOC9KN0xjUWd1RldlTDZ0ZmwyMUxCSFhBTUN4NURITnVuNHdKM3kxTGdOYSIsIkJpbyI6Ik1hZWNlbmFzIGxlbyBvZGlvLCBjb25kaW1lbnR1bSBpZCwgbHVjdHVzIG5lYywgbW9sZXN0aWUgc2VkLCBqdXN0by4gUGVsbGVudGVzcXVlIHZpdmVycmEgcGVkZSBhYyBkaWFtLiBDcmFzIHBlbGxlbnRlc3F1ZSB2b2x1dHBhdCBkdWkuXG5cbk1hZWNlbmFzIHRyaXN0aXF1ZSwgZXN0IGV0IHRlbXB1cyBzZW1wZXIsIGVzdCBxdWFtIHBoYXJldHJhIG1hZ25hLCBhYyBjb25zZXF1YXQgbWV0dXMgc2FwaWVuIHV0IG51bmMuIFZlc3RpYnVsdW0gYW50ZSBpcHN1bSBwcmltaXMgaW4gZmF1Y2lidXMgb3JjaSBsdWN0dXMgZXQgdWx0cmljZXMgcG9zdWVyZSBjdWJpbGlhIEN1cmFlOyBNYXVyaXMgdml2ZXJyYSBkaWFtIHZpdGFlIHF1YW0uIFN1c3BlbmRpc3NlIHBvdGVudGkuIiwiUGhvbmUiOiIzMTU0OTkzMDYxIiwiUHJvZmlsZUltYWdlIjoiY2p5anh1ZGJ3MDAwMGMxcWc5bXNydDZ2dSIsIldhbGxJbWFnZSI6ImNqeWp4dWRidzAwMDBjMXFnOW1zcnQ2dnUiLCJBZGRyZXNzSUQiOjAsIkFzc2V0cyI6bnVsbCwiVmVuZG9yIjpudWxsLCJDcmVhdGVkQXQiOiIyMDE5LTA4LTI4VDA4OjU4OjI2LjEwMzU4NVoiLCJVcGRhdGVkQXQiOiIyMDE5LTA4LTI4VDA4OjU4OjI2LjEwMzU4NVoifX0.KIegCNeVBfpe7ghxoFaOER-w1iKsM2CFHBdTqFZvPW8",
			Error: false,
		},
		{
			Title: "should return an error if Decode token fails",
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJsb29rYmpdHoiLCJhdWQiOiJsb2NhbGhvc3Q6MzAwMSIsImV4cCI6MTU5ODM3MDg4OCwibmJmIjoxNTY3MjY4Njg4LCJpYXQiOjE1NjcyNjY4ODgsIlVzZXIiOnsiSUQiOiJjank3bjh3MTEwMDAxbGhxZ2gzNmY2ZHRsIiwiRmlyc3RuYW1lIjoiQWxldGEiLCJMYXN0bmFtZSI6IkNlY2VyZSIsIlVzZXJuYW1lIjoiYWNlY2VyZTAiLCJFbWFpbCI6ImFjZWNlcmUwQHJlZmVyZW5jZS5jb20iLCJQYXNzd29yZCI6IiQyYSQxNCQzZXZHbDdWUWtyOC9KN0xjUWd1RldlTDZ0ZmwyMUxCSFhBTUN4NURITnVuNHdKM3kxTGdOYSIsIkJpbyI6Ik1hZWNlbmFzIGxlbyBvZGlvLCBjb25kaW1lbnR1bSBpZCwgbHVjdHVzIG5lYywgbW9sZXN0aWUgc2VkLCBqdXN0by4gUGVsbGVudGVzcXVlIHZpdmVycmEgcGVkZSBhYyBkaWFtLiBDcmFzIHBlbGxlbnRlc3F1ZSB2b2x1dHBhdCBkdWkuXG5cbk1hZWNlbmFzIHRyaXN0aXF1ZSwgZXN0IGV0IHRlbXB1cyBzZW1wZXIsIGVzdCBxdWFtIHBoYXJldHJhIG1hZ25hLCBhYyBjb25zZXF1YXQgbWV0dXMgc2FwaWVuIHV0IG51bmMuIFZlc3RpYnVsdW0gYW50ZSBpcHN1bSBwcmltaXMgaW4gZmF1Y2lidXMgb3JjaSBsdWN0dXMgZXQgdWx0cmljZXMgcG9zdWVyZSBjdWJpbGlhIEN1cmFlOyBNYXVyaXMgdml2ZXJyYSBkaWFtIHZpdGFlIHF1YW0uIFN1c3BlbmRpc3NlIHBvdGVudGkuIiwiUGhvbmUiOiIzMTU0OTkzMDYxIiwiUHJvZmlsZUltYWdlIjoiY2p5anh1ZGJ3MDAwMGMxcWc5bXNydDZ2dSIsIldhbGxJbWFnZSI6ImNqeWp4dWRidzAwMDBjMXFnOW1zcnQ2dnUiLCJBZGRyZXNzSUQiOjAsIkFzc2V0cyI6bnVsbCwiVmVuZG9yIjpudWxsLCJDcmVhdGVkQXQiOiIyMDE5LTA4LTI4VDA4OjU4OjI2LjEwMzU4NVoiLCJVcGRhdGVkQXQiOiIyMDE5LTA4LTI4VDA4OjU4OjI2LjEwMzU4NVoifX0.KIegCNeVBfpe7ghxoFaOER-w1iKsM2CFHBdTqFZvPW8",
			Error: true,
		},
	}

	for _, test := range testCases {
		payload, err := JWT.DecodeToken(test.Token)

		if test.Error {
			assert.NotNil(t, err, test.Title)
		} else {
			assert.Nil(t, err)
			assert.IsType(t, payload, models.User{}, test.Title)
		}
	}
}

func TestHashPassword(t *testing.T) {
	testCases := []struct {
		Title    string
		Password string
	}{
		{
			Title:    "Should return hashed password when HashPassword is called",
			Password: "test1234",
		},
		{
			Title:    "Should return hashed password when HashPassword is called",
			Password: "another pass",
		},
	}

	for _, test := range testCases {
		password, err := JWT.HashPassword(test.Password)

		assert.Nil(t, err)
		assert.True(t, len(password) > 10)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	testCases := []struct {
		Title    string
		Password string
		Hash     string
		Result   bool
	}{
		{
			Title:    "Should return true when CheckPasswordHash is called with valid password test1234",
			Password: "test1234",
			Hash:     "$2a$10$FZlR0DJUeYts8Io2YL0SJe0tjh4F6PqsUmyRGG1nzrAusRyzRVhBK",
		},
		{
			Title:    "Should return true when CheckPasswordHash is called with valid password another pass",
			Password: "another pass",
			Hash:     "$2a$10$UZ2A7ESjIbw5daDdi77bLed8vkSXLRug1PSOdja0DO.A79Sssvj8CK",
		},
	}

	for _, test := range testCases {
		isValid := JWT.CheckPasswordHash(test.Password, test.Hash)

		assert.True(t, isValid, test.Title)
	}
}
