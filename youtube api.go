package main

import (
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

// Download the youtube video as a mp4
func songDownload(songID string) {
	client := youtube.Client{}

	video, err := client.GetVideo(songID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	songTitle := video.Title + ".mp4"

	file, err := os.Create(songTitle)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}