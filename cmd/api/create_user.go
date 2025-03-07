package main

import (
	"app/internal/domain/entity"
	"app/internal/infrastructure/configuration"
	"app/internal/infrastructure/persistence"
	"app/internal/infrastructure/repository"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Wrong number of arguments: must be <email> and <password> only")
		return
	}
	email := strings.TrimSpace(args[1])
	password := strings.TrimSpace(args[2])

	if !isValidEmail(email) {
		fmt.Println("Invalid email format")
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
