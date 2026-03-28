package commands

import (
	"errors"
	"moshi/internal/texts"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func getAdequatedMediaDir(filetype string) string {
	filetype = strings.ToLower(filetype)
	userOS := runtime.GOOS

	if filetype == "mp3" {
		return "Music"
	} // case is audio

	// case video:
	switch userOS {
		case "windows":
			return "Videos"
		case "linux":
			return "Videos"
		case "android", "darwin":
			return "Movies"
		default:
			return "Movies"
	}
}

func GetMediaDir(filetype string) (*string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return nil, errors.New(texts.PortugueseErrors.AccessError)
	}

	userOS := runtime.GOOS
	

	typeDir := getAdequatedMediaDir(filetype)

	homeDir := ""

	switch userOS {
	case "windows", "linux", "darwin":
		homeDir = filepath.Join(home, typeDir)
	case "android":
		homeDir = filepath.Join(home, "storage", "shared", typeDir)
	default:
		return nil, errors.New(texts.PortugueseErrors.AccessError)
	}

	return &homeDir, nil
}
