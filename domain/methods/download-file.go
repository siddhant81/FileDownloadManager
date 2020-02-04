package methods

import (
	"io"
	"net/http"
	"os"
	"path"
)

func DownloadFile(url string) (string, error){
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "URL not working", err
	}
	defer resp.Body.Close()

	// Create the file
	filepath := path.Base(resp.Request.URL.String())
	out, err := os.Create(filepath)
	if err != nil {
		return "File creation error", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "File writing error", err
	}

	return filepath, nil
}