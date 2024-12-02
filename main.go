package main

import (
	"fmt"
)

func main() {
	var search string
	fmt.Println("Song search: ")
	fmt.Scan(&search)

	videoID := "auM5KOAlQlg"
	songDownload(videoID)
}