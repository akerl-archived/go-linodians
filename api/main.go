package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// TODO: photo support

const (
	dataURL        = "https://www.linode.com/about"
	photoURLFormat = "https://www.linode.com/media/images/employees/%s.png"
)

// Employee describes a human at Linode
type Employee struct {
	Username string
	Fullname string
	Title    string
	Social   map[string]string
}

// Load employees from site
func Load() ([]Employee, error) {
	page, err := download(dataURL)
	if err != nil {
		return []Employee{}, err
	}

	reader := bytes.NewReader(page)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return []Employee{}, err
	}

	divs := doc.Find("div.employee-display")
	for _, node := range divs.Nodes {
		fmt.Printf("%+v\n", node)
	}
	return []Employee{}, nil
}

// Photo downloads the photo for an employee
func (e Employee) Photo() ([]byte, error) {
	if e.Username == "" {
		return []byte{}, fmt.Errorf("Username not set")
	}
	url := fmt.Sprintf(photoURLFormat, e.Username)
	return download(url)
}

func download(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return []byte{}, fmt.Errorf("HTTP call failed: %s (%d)", url, res.StatusCode)
	}
	return ioutil.ReadAll(res.Body)
}
