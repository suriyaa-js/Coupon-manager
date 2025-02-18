package main

import (
	"fmt"

	"suriyaa.com/coupon-manager/server"
)

const (
	configPath = "./env/application.yaml" // Path to the configuration file, Can be set as an environment variable
)

func main() {
	// Load configuration
	config, err := server.LoadConfig(configPath)
	if err != nil {
		fmt.Println("Error loading configuration: ", err)
		panic(err)
	}

	sc := server.NewServer(&config)
	sc.ConfigureAPI(&config)

	err = sc.Serve()
	if err != nil {
		fmt.Println("Error starting server: ", err)
		panic(err)
	}

}
