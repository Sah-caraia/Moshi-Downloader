package texts

import "moshi/internal/models"

var EnglishErrors = models.ErrText{
	NotFoundLibrary: "The program \"yt-dlp\" is not installed! To install it, first install <Python>. Then run \"pip install -U yt-dlp\" and try again.",
	NetworkError: "It looks like you're not connected to the internet :/\nPlease try again later.",
	AccessError: "The file could not be saved due to a permission error.",
	DownloadError: "An error occurred! The media might not be available.",
}

var EnglishSuccess = models.SusText{
	MediaInstalled: "Media downloaded successfully! ✓",
	AllSuccess: "All media were downloaded successfully! ✓",
}
