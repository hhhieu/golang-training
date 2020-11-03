package database

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// MySQLUserName presents the name of environment vairable that saves the MySQL user name
	MySQLUserName = "MYSQL_USER_NAME"
	// MySQLPassword presents the name of environment vairable that saves the MySQL password
	MySQLPassword = "MYSQL_PASSWORD"
	// MySQLAddress presents the name of environment vairable that saves the MySQL address. Default is "localhost"
	MySQLAddress = "MYSQL_ADDRESS"
	// MySQLPort presents the name of environment vairable that saves the MySQL port. Default is 3306
	MySQLPort = "MYSQL_PORT"
	// MySQLDatabase presents the name of environment vairable that saves the MySQL database
	MySQLDatabase = "MYSQL_DATABASE"
)

const defaultMySQLPort = 3306
const defaultMySQLAddress = "localhost"

// Config contains the data source name(DSN)
type Config struct {
	MYSQLDSN string
}

// GetEnvironmentVariableNames get name of all environment variable which is supported.
func GetEnvironmentVariableNames() []string {
	return []string{
		MySQLAddress,
		MySQLDatabase,
		MySQLPassword,
		MySQLPort,
		MySQLUserName,
	}
}

// LoadConfig reads the environment variables then generates the data source name(DSN)
func LoadConfig() (*Config, error) {
	username, exist := os.LookupEnv(MySQLUserName)
	if !exist {
		return nil, fmt.Errorf("Environment variable %s is not set", MySQLUserName)
	}
	password, exist := os.LookupEnv(MySQLPassword)
	if !exist {
		return nil, fmt.Errorf("Environment variable %s is not set", MySQLPassword)
	}
	database, exist := os.LookupEnv(MySQLDatabase)
	if !exist {
		return nil, fmt.Errorf("Environment variable %s is not set", MySQLDatabase)
	}
	address, exist := os.LookupEnv(MySQLAddress)
	if !exist {
		address = defaultMySQLAddress
	}
	portStr, exist := os.LookupEnv(MySQLPort)
	port, err := strconv.ParseUint(portStr, 10, 32)
	if !exist || err != nil {
		port = defaultMySQLPort
	}
	return &Config{
		MYSQLDSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, database),
	}, nil
}
