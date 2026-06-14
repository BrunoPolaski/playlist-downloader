package youtube

import (
	"fmt"
	"os"
	"os/exec"
)

type Downloader struct {
	OutputDir string
}

func New(outputDir string) *Downloader {
	return &Downloader{
		OutputDir: outputDir,
	}
}

func (d *Downloader) DownloadPlaylist(playlistURL string) error {
	if err := os.MkdirAll(d.OutputDir, 0755); err != nil {
		return err
	}

	args := []string{
		"--yes-playlist",
		"--extract-audio",
		"--audio-format", "mp3",
		"--audio-quality", "0",
		"--embed-metadata",
		"--embed-thumbnail",
		"-o",
		fmt.Sprintf("%s/%%(playlist_index)s - %%(title)s.%%(ext)s", d.OutputDir),
		playlistURL,
	}

	cmd := exec.Command("yt-dlp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
