package main

import (
	config "final-project/config"
	repo "final-project/repo"
	router "final-project/router"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	r := gin.Default()
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config file: ", err)
	}

	c := []string{"postgres://", config.POSTGRES_USER, ":", config.POSTGRES_PASSWORD, "@postgres", ":", config.POSTGRES_PORT, "/", config.POSTGRES_DB, "?sslmode=disable"}

	var builder strings.Builder
	for _, item := range c {
		builder.WriteString(item)
	}

	dbSource := builder.String()

	repo.ConnectDatabase(config)
	runDBMigration(config.MIGRATION_URL, dbSource)
	router.InitRouter(r)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance: ", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up ", err)
	}

	log.Println("db migrated successfully")
}
