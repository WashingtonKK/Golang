package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//Struct to hold the data from the yaml file
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Server struct {
		Port int `yaml:"port"`
	}
}

func main() {
	//Read the config file
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	//Parse the config file.
	config := Config{}
	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		log.Fatalf("Failed to parse config file: %V", err)
	}

	//Print values from the config file
	fmt.Printf("Database host: %s\n", config.Database.Host)
	fmt.Printf("Database port: %d\n", config.Database.Port)
	fmt.Printf("Database username: %s\n", config.Database.Username)
	fmt.Printf("Database password: %s\n", config.Database.Password)
	fmt.Printf("Server port: %d\n", config.Server.Port)
}
