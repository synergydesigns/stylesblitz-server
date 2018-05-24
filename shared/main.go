package main

import (
	"fmt"

	"gitlab.com/synergy-designs/style-blitz/shared/service"
)

func main() {
	user, err := service.GetUserByID(1)
	if err == nil {
		fmt.Println(user)
	}
}
