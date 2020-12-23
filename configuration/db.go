package configuration

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct{
	DbUser string
	DbPass string
	DbName string
}

func  SqlConfig(config *Configuration) (*sql.DB, error) {

	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s",
		config.DbUser, config.DbName, config.DbPass))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (c *Configuration) ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}

//NewConfiguration read file, return configuration
func NewConfiguration(path string) *Configuration {
	var configuration Configuration
	configuration.ReadFile(path)

	return &configuration
}
