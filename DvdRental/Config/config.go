package config

import (
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct{
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string  `yaml:"user"` 
		Password string `yaml:"password"`
		Database string `yaml:"dbname"`
	} `yaml:"database"`
}

func loadConfig() Config{
	file, err := os.ReadFile("Config/config.yaml")

	if err != nil {
		log.Println("Error reading config file:", err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error unmarshalling config file:", err)
		
	} else {
		log.Println("Config loaded successfully")
	}	

	return config
}
