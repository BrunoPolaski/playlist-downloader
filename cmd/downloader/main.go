package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/brunopolaski/youtube-playlist-downloader/internal/youtube"
)

func main() {
	playlist := flag.String("playlist", "", "YouTube playlist URL")
	output := flag.String("output", "./downloads", "Output directory")

	flag.Parse()

	if *playlist == "" {
		log.Fatal("playlist URL is required")
	}

	downloader := youtube.New(*output)

	fmt.Println("Downloading playlist...")

	if err := downloader.DownloadPlaylist(*playlist); err != nil {
		log.Fatalf("download failed: %v", err)
	}

	fmt.Println("Done.")
}
