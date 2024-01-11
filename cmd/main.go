package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/raphael251/users-crud/configs"
	"github.com/raphael251/users-crud/internal/infra/database"
	"github.com/raphael251/users-crud/internal/infra/web/handlers"
	"github.com/raphael251/users-crud/internal/infra/web/server"
)

func main() {
	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName,
	))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	server := server.NewServer(configs.ServerPort)

	userRepository := database.NewUserRepository(db)

	applicationHandler := handlers.NewApplicationHandler()
	userHandler := handlers.NewUserHandler(userRepository)

	server.AddHandler("/health", applicationHandler.Health)
	server.AddHandler("/users", userHandler.Create)

	server.Start()
}
