package util

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

// Fetch is the equivalent of curl -sSL or wget to simply fetch a file
// from the Web. The first argument is always the URL. The second
// argument is the path to local file to which to write the downloaded
// file. Note that if the file exists it will be overwritten. The third
// argument is the maximum number of seconds to wait for a response.
//
func Fetch(url, local string, timeout time.Duration) error {
	var err error
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Second*timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	file, err := os.Create(local)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	return nil
}
