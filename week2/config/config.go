package config

type Config struct {
	MSQLDNS string
}

func NewDefaultConfig() *Config {
	return &Config{
		MSQLDNS: "hieu:password@tcp(127.0.0.1:3306)/dogfood?charset=utf8mb4&parseTime=True&loc=Local",
	}
}
