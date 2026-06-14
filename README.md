## Install dependencies:

#### Arch Linux
```sh
sudo pacman -S ffmpeg
pip install yt-dlp

# or

yay -S yt-dlp
```

#### Run:

```sh
go run ./cmd/downloader \
  --playlist "https://www.youtube.com/playlist?list=XXXX"
```

#### Output:

```sh
downloads/
├── 001 - My First Video.mp3
├── 002 - Another Video.mp3
├── 003 - ...
```
