package main

import (
	"app/internal/domain/entity"
	"app/internal/domain/valueobject"
	"app/internal/infrastructure/configuration"
	"app/internal/infrastructure/persistence"
	"app/internal/infrastructure/repository"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Wrong number of arguments: must be <email> and <password> only")
		return
	}
	email, err := valueobject.NewEmail(strings.TrimSpace(args[1]))
	if err != nil {
		fmt.Println("Invalid email format")
		return
	}
	password, err := valueobject.NewPassword(strings.TrimSpace(args[2]))
	if err != nil {
		fmt.Println("Invalid password format")
		return
	}
	cfg := configuration.LoadConfig()
	db := persistence.ConnectDatabase(cfg)
	repo := repository.NewUserPgRepository(db)
	user, err := entity.NewUser(email, password)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	err = repo.Create(user)
	if err != nil {
		fmt.Println("Failed to create user in database:", err)
		return
	}
	fmt.Println(user)
	fmt.Println("User created successfully:", user.Email)
	return
}
