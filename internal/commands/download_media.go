package commands

import (
	"errors"
	"moshi/internal/texts"
	"os"
	"os/exec"
	"path/filepath"
)

func DownloadMedia(url string, filetype string) (error) {
	var err error = nil

	// check yt-dlp lib
	_, err = exec.LookPath("yt-dlp")
	if err != nil { // case is not installed
		return errors.New(texts.PortugueseErrors.NotFoundLibrary)
	}

	mediaDir, err := GetMediaDir(filetype)
	if err != nil {
		return err
	}

	filename := "%(title)s." + filetype // to get filename
	exitPath := filepath.Join(*mediaDir, filename) // join all: mediapath + filename
	err = os.MkdirAll(exitPath, 0755)
	if err != nil {
		return errors.New(texts.PortugueseErrors.AccessError)
	}

	exit, err := exec.Command(
		"yt-dlp",
		"-t", filetype,
		"-o", exitPath,
		url,
	).CombinedOutput()
	if err != nil {
		return errors.New(string(exit))
	}

	return nil
}
