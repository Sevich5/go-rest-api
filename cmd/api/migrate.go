package main

import (
	"app/internal/infrastructure/configuration"
	"app/internal/infrastructure/persistence/model"
	"ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		cfg := configuration.LoadConfig()
		dbName := cfg.Database.Database
		switch args[1] {
		case "dsn":
			{
				break
			}
		case "dsn_dev":
			{
				dbName = cfg.Database.DevDatabase
				break
			}
		default:
			return
		}
		out := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?search_path=public&sslmode=disable",
			cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, dbName,
		)
		_, err := io.WriteString(
			os.Stdout, out,
		)
		if err != nil {
			return
		}
	} else {
		stmts, err := gormschema.New("postgres").Load(
			&model.User{},
		)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		_, err = io.WriteString(os.Stdout, stmts)
		if err != nil {
			return
		}
	}
}
