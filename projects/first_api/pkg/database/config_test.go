package database

import (
	"os"
	"reflect"
	"testing"
)

func unsetEnvironmentVairables() {
	envVarList := GetEnvironmentVariableNames()
	for _, v := range envVarList {
		os.Unsetenv(v)
	}
}

func setEnvironmentVairables(t *testing.T) {
	envVarList := GetEnvironmentVariableNames()
	for _, v := range envVarList {
		switch v {
		case MySQLAddress:
			os.Setenv(v, "address")
		case MySQLDatabase:
			os.Setenv(v, "database")
		case MySQLPassword:
			os.Setenv(v, "password")
		case MySQLPort:
			os.Setenv(v, "80")
		case MySQLUserName:
			os.Setenv(v, "username")
		default:
			t.Errorf("Unknown environment variable %s", v)
		}
	}
}

func TestLoadConfigSuccessful(t *testing.T) {
	t.Run("Load config successful", func(t *testing.T) {
		// Set environment variables
		setEnvironmentVairables(t)
		// Unset environment variables
		defer unsetEnvironmentVairables()
		wantedMySQLDSN := "username:password@tcp(address:80)/database?charset=utf8mb4&parseTime=True&loc=Local"
		config, err := LoadConfig()
		if err != nil {
			t.Errorf("LoadConfig() error = %v, don't want error", err)
		} else if config.MYSQLDSN != wantedMySQLDSN {
			t.Errorf("LoadConfig() = %v, want %v", config.MYSQLDSN, wantedMySQLDSN)
		}
	})
}

func TestLoadConfigUnsetAddress(t *testing.T) {
	t.Run("Load config with unsetting port", func(t *testing.T) {
		// Set environment variables
		setEnvironmentVairables(t)
		// Unset environment variables
		defer unsetEnvironmentVairables()
		// Unset address
		os.Unsetenv(MySQLAddress)
		wantedMySQLDSN := "username:password@tcp(localhost:80)/database?charset=utf8mb4&parseTime=True&loc=Local"
		config, err := LoadConfig()
		if err != nil {
			t.Errorf("LoadConfig() error = %v, don't want error", err)
		} else if config.MYSQLDSN != wantedMySQLDSN {
			t.Errorf("LoadConfig() = %v, want %v", config.MYSQLDSN, wantedMySQLDSN)
		}
	})
}

func TestLoadConfigUnsetPort(t *testing.T) {
	t.Run("Load config with unsetting port", func(t *testing.T) {
		// Set environment variables
		setEnvironmentVairables(t)
		// Unset environment variables
		defer unsetEnvironmentVairables()
		// Unset port
		os.Unsetenv(MySQLPort)
		wantedMySQLDSN := "username:password@tcp(address:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
		config, err := LoadConfig()
		if err != nil {
			t.Errorf("LoadConfig() error = %v, don't want error", err)
		} else if config.MYSQLDSN != wantedMySQLDSN {
			t.Errorf("LoadConfig() = %v, want %v", config.MYSQLDSN, wantedMySQLDSN)
		}
	})
}

func TestLoadConfigUnsetDatabase(t *testing.T) {
	t.Run("Load config with unsetting database", func(t *testing.T) {
		// Set environment variables
		setEnvironmentVairables(t)
		// Unset environment variables
		defer unsetEnvironmentVairables()
		// Unset database
		os.Unsetenv(MySQLDatabase)
		_, err := LoadConfig()
		if err == nil {
			t.Errorf("LoadConfig() no error, want error")
		}
	})
}

func TestLoadConfigUnsetPassword(t *testing.T) {
	t.Run("Load config with unsetting password", func(t *testing.T) {
		// Set environment variables
		setEnvironmentVairables(t)
		// Unset environment variables
		defer unsetEnvironmentVairables()
		// Unset password
		os.Unsetenv(MySQLPassword)
		_, err := LoadConfig()
		if err == nil {
			t.Errorf("LoadConfig() no error, want error")
		}
	})
}

func TestLoadConfigUnsetUserName(t *testing.T) {
	t.Run("Load config with unsetting user name", func(t *testing.T) {
		// Set environment variables
		setEnvironmentVairables(t)
		// Unset environment variables
		defer unsetEnvironmentVairables()
		// Unset user name
		os.Unsetenv(MySQLUserName)
		_, err := LoadConfig()
		if err == nil {
			t.Errorf("LoadConfig() no error, want error")
		}
	})
}

func TestGetEnvironmentVariableNames(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Get environment variable names",
			want: []string{
				MySQLAddress,
				MySQLDatabase,
				MySQLPassword,
				MySQLPort,
				MySQLUserName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvironmentVariableNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvironmentVariableNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
