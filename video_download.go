package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {
	// YouTube video URL
	videoURL := "https://www.youtube.com/watch?v=VIDEO_ID"

	// Create a new YouTube client
	client := youtube.Client{}

	// Get video information
	video, err := client.GetVideo(videoURL)
	if err != nil {
		log.Fatalf("Error getting video info: %v", err)
	}

	// Print video title
	fmt.Printf("Downloading video: %s\n", video.Title)

	// Get the highest resolution format available
	formats := video.Formats.WithAudioChannels() // Ensure the format includes audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		log.Fatalf("Error getting video stream: %v", err)
	}
	defer stream.Close()

	// Create a new file to save the video
	file, err := os.Create(video.Title + ".mp4")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	// Copy the video stream to the file
	_, err = file.ReadFrom(stream)
	if err != nil {
		log.Fatalf("Error saving video: %v", err)
	}

	fmt.Println("Video downloaded successfully!")
}