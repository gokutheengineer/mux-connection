package main

import (
	"fmt"
	"log"

	muxgo "github.com/muxinc/mux-go"
	"github.com/spf13/viper"
)

func main() {

	initConfig()

	// Get values from the config
	muxTokenID := viper.GetString("MUX_TOKEN_ID")
	muxTokenSecret := viper.GetString("MUX_TOKEN_SECRET")

	fmt.Println("Mux Token ID: ", muxTokenID)
	fmt.Println("Mux Token Secret: ", muxTokenSecret)

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(muxTokenID, muxTokenSecret),
		))
	// Create the Asset
	asset, err := client.AssetsApi.CreateAsset(muxgo.CreateAssetRequest{
		Input: []muxgo.InputSettings{
			muxgo.InputSettings{
				Url: "https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4",
			},
		},
		PlaybackPolicy: []muxgo.PlaybackPolicy{muxgo.PUBLIC},
	})

	// Check everything was good, and output the playback URL
	if err == nil {
		fmt.Printf("Playback URL: https://stream.mux.com/%s.m3u8 \n", asset.Data.PlaybackIds[0].Id)
	} else {
		fmt.Printf("Oh no, there was an error: %s \n", err)
	}
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
