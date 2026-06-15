package youtube

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

type Downloader struct {
	OutputDir string
	Workers   int
}

func New(outputDir string, workers int) *Downloader {
	return &Downloader{
		OutputDir: outputDir,
		Workers:   workers,
	}
}

func (d *Downloader) DownloadPlaylist(playlistURL string) error {
	if err := os.MkdirAll(d.OutputDir, 0755); err != nil {
		return err
	}

	videoURLs, err := d.getPlaylistVideos(playlistURL)
	if err != nil {
		return err
	}

	fmt.Printf("Found %d videos\n", len(videoURLs))
	fmt.Printf("Starting %d workers\n", d.Workers)

	jobs := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= d.Workers; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for url := range jobs {
				fmt.Printf("[Worker %d] Downloading %s\n", workerID, url)

				if err := d.downloadVideo(url); err != nil {
					fmt.Printf("[Worker %d] ERROR: %v\n", workerID, err)
					continue
				}

				fmt.Printf("[Worker %d] Done\n", workerID)
			}
		}(i)
	}

	for _, url := range videoURLs {
		jobs <- url
	}

	close(jobs)
	wg.Wait()

	return nil
}

func (d *Downloader) getPlaylistVideos(playlistURL string) ([]string, error) {
	cmd := exec.Command(
		"yt-dlp",
		"--flat-playlist",
		"--print",
		"url",
		playlistURL,
	)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var urls []string

	scanner := bufio.NewScanner(bytes.NewReader(output))

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	return urls, scanner.Err()
}

func (d *Downloader) downloadVideo(videoURL string) error {
	cmd := exec.Command(
		"yt-dlp",
		"--extract-audio",
		"--audio-format", "mp3",
		"--audio-quality", "0",
		"--embed-metadata",
		"--embed-thumbnail",
		"-o",
		fmt.Sprintf("%s/%%(title)s.%%(ext)s", d.OutputDir),
		videoURL,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
