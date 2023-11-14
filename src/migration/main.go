package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {

	_ = InitConfig()

	db, err := sql.Open(viper.GetString("DB_DRIVER"), viper.GetString("DB_SOURCE_MASTER"))
	if err != nil {
		log.Panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(viper.GetString("DB_MIGRATE_PATH"), viper.GetString("DB_DRIVER"), driver)
	if err != nil {
		log.Panic(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func InitConfig() error {
	viper.SetConfigFile(`.env`)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error load .env")
	}
	return nil
}
