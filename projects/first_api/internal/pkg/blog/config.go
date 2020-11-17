package blog

import (
	"fmt"

	"github.com/hhhieu/golang-training/first_api/pkg/database"
	"github.com/hhhieu/golang-training/first_api/pkg/system"
	"gopkg.in/yaml.v2"
)

const defaultURL = "localhost:3306"

// WebConfig contains the web configuration
type WebConfig struct {
	Address string
}

// Config contains all configurations
type Config struct {
	Database database.Config
	Web      WebConfig
}

// LoadConfig load the configuration from file
func LoadConfig(ioUtil system.IOUtiler, filePath string) (*Config, error) {
	// Read the configuration file
	configBytes, err := ioUtil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// Parse the yaml data to struct
	config := Config{
		Database: database.Config{
			URL: defaultURL,
		},
	}
	yaml.Unmarshal(configBytes, &config)
	// Validate the database configuration
	if err := config.Database.Validate(); err != nil {
		return nil, fmt.Errorf("Loading database configuration failed with error %v", err)
	}
	return &config, nil
}
