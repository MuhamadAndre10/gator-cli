package config

import (
	"encoding/json"
	"os"
)

func getConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.gatorconfig.json"
}

type Config struct {
	DB_URL       string `json:"db_url"`
	CURRENT_USER string `json:"current_user"`
}

func Read() (*Config, error) {

	jsonFile, err := os.ReadFile(getConfigPath())
	if err != nil {
		return nil, err
	}

	var config Config
	json.Unmarshal(jsonFile, &config)
	return &config, nil
}

func (c *Config) SetUser(user string) {

	// set user. now on json file don't have current user value.
	c.CURRENT_USER = user
	Write(c)
}

func (c *Config) SetDBUrl(dbUrl string) {

	// set db url. now on json file don't have db url value.
	c.DB_URL = dbUrl
	Write(c)
}

func (c *Config) GetUser() string {
	return c.CURRENT_USER
}

func Write(cfg *Config) {

	cfgJson, err := json.Marshal(cfg)
	if err != nil {
		return
	}
	os.WriteFile(getConfigPath(), cfgJson, 0644)

}
