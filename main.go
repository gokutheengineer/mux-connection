package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	initConfig()

	// Get values from the config
	muxTokenID := viper.GetString("MUX_TOKEN_ID")
	muxTokenSecret := viper.GetString("MUX_TOKEN_SECRET")
	creatorUserID := viper.GetString("CREATOR_USER_ID")

	fmt.Println("Mux Token ID: ", muxTokenID)
	fmt.Println("Mux Token Secret: ", muxTokenSecret)

	UploadAndNotify(muxTokenID, muxTokenSecret, creatorUserID)
}

func initConfig() {
	// Set the file name of the .env file
	viper.SetConfigFile("config.env")

	// Read the .env file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	// Automatically override environment variables if they exist
	viper.AutomaticEnv()
}
