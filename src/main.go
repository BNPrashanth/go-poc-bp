package main

import (
	"fmt"

	c "./config"
	"github.com/spf13/viper"
)

func main() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	
	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.Database.DBName)
	fmt.Println("Port is\t\t", configuration.Server.Port)
	fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
	fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Database is\t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))
	fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
	fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_VAR"))

}
