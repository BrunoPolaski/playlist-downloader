package main

import (
	"flag"
	"log"

	"github.com/BrunoPolaski/playlist-downloader/internal/youtube"
)

func main() {
	playlist := flag.String("playlist", "", "playlist url")
	output := flag.String("output", "./downloads", "output directory")
	workers := flag.Int("workers", 5, "number of concurrent downloads")

	flag.Parse()

	if *playlist == "" {
		log.Fatal("playlist is required")
	}

	d := youtube.New(*output, *workers)

	if err := d.DownloadPlaylist(*playlist); err != nil {
		log.Fatal(err)
	}
}
