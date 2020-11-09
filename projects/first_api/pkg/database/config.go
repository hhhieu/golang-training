package database

import "fmt"

// Config configure the information of connection database
type Config struct {
	Type     string
	URL      string
	Username string
	Password string
	DBName   string
}

// Validate the properties of configuration
func (C Config) Validate() error {
	// Validate that all properties are not empty
	if C.Type == "" {
		return fmt.Errorf("Type is empty")
	} else if C.URL == "" {
		return fmt.Errorf("URL is empty")
	} else if C.Username == "" {
		return fmt.Errorf("Username is empty")
	} else if C.Password == "" {
		return fmt.Errorf("Password is empty")
	} else if C.DBName == "" {
		return fmt.Errorf("DBName is empty")
	}
	// Validate the database type
	isSupport := false
	for _, t := range GetSupportedType() {
		if t == C.Type {
			isSupport = true
			break
		}
	}
	if !isSupport {
		return fmt.Errorf("Database type %v is not supported", C.Type)
	}

	return nil
}
