package database

import (
	"fmt"
)

type Config struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     string
}

func GetConnectionString(config Config) string {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.Port, config.User, config.Name, config.Password)
	return connectionString
}

func UrlToConfig(url string) Config {
	var config Config
	fmt.Sscanf(url, "postgres://%s:%s@%s:%s/%s", &config.User, &config.Password, &config.Host, &config.Port, &config.Name)
	return config
}
