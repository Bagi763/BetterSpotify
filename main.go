package main

import (
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

// ExampleDownload : Example code for how to use this package for download video.
func main() {
	videoID := "auM5KOAlQlg"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}