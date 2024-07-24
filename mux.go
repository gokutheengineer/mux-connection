package main

import (
	"fmt"
	"log"

	muxgo "github.com/muxinc/mux-go"
)

// Function to upload a video and notify the application
func uploadAndNotify(muxTokenID, muxTokenSecret, uploaderUserID string) {
	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(muxTokenID, muxTokenSecret),
		))

	// Create the Asset
	asset, err := client.AssetsApi.CreateAsset(muxgo.CreateAssetRequest{
		Input: []muxgo.InputSettings{
			{
				Url: "https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4",
			},
		},
		PlaybackPolicy: []muxgo.PlaybackPolicy{muxgo.PUBLIC},
	})

	// Check everything was good, and output the playback URL
	if err == nil {
		playbackURL := fmt.Sprintf("https://stream.mux.com/%s.m3u8", asset.Data.PlaybackIds[0].Id)
		fmt.Printf("Playback URL: %s\n", playbackURL)

		// Notify the app and grant the uploader the "creator" role
		notifyAppAndGrantRole(uploaderUserID, asset.Data.Id)
	} else {
		fmt.Printf("Oh no, there was an error: %s \n", err)
	}
}

// MOCK
// Function to notify the app and grant the uploader the "creator" role
func notifyAppAndGrantRole(userID, assetID string) {
	err := grantCreatorRole(userID)
	if err != nil {
		log.Fatalf("Failed to grant creator role: %s", err)
	}

	fmt.Printf("Granted creator role to user: %s for asset: %s\n", userID, assetID)
}

// MOCK
func grantCreatorRole(userID string) error {
	fmt.Printf("User %s has been granted the creator role.\n", userID)
	return nil
}

// Function to get playback URL for a watcher
func getPlaybackURL(assetID string, userID string) (string, error) {
	// Check if the user has the watcher role
	if !userHasRole(userID, "watcher") {
		return "", fmt.Errorf("user does not have watcher role")
	}

	// Retrieve the playback URL from your database or storage
	playbackURL := getPlaybackURLFromDatabase(assetID)

	return playbackURL, nil
}

// Mock function to check if a user has a specific role
func userHasRole(_, _ string) bool {
	// Mock implementation
	return true
}

// Mock function to get playback URL from a database
func getPlaybackURLFromDatabase(assetID string) string {
	// Mock implementation
	return fmt.Sprint("https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4")
}
