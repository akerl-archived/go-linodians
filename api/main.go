package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	dataURL        = "https://www.linode.com/about"
	photoURLFormat = "https://www.linode.com/media/images/employees/%s.png"
)

func download(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		if res.Body.Close() != nil {
			panic(err)
		}
	}()
	if res.StatusCode != 200 {
		return []byte{}, fmt.Errorf("HTTP call failed: %s (%d)", url, res.StatusCode)
	}
	return ioutil.ReadAll(res.Body)
}
