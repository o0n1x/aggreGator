package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

var testconfigFilePath = ""

type Config struct {
	DB_URL          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	cnfgPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(cnfgPath)
	if err != nil {
		return Config{}, err
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(newuser string) error {
	c.CurrentUserName = newuser
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func getConfigPath() (string, error) {
	if testconfigFilePath != "" {
		return testconfigFilePath, nil
	}
	homepath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homepath, configFileName), nil
}

func write(cnfg Config) error {
	data, err := json.Marshal(cnfg)
	if err != nil {
		return err
	}
	cnfgPath, err := getConfigPath()
	if err != nil {
		return err
	}
	err = os.WriteFile(cnfgPath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
