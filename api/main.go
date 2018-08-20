package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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
	el := make([]Employee, len(divs.Nodes))

	divs.Each(func(i int, s *goquery.Selection) {
		social := make(map[string]string)
		s.Find("a.employee-link").Each(func(_ int, ss *goquery.Selection) {
			class := ss.AttrOr("class", "")
			parts := strings.Split(class, "-")
			site := parts[len(parts)-1]
			link := ss.AttrOr("href", "")
			social[site] = link
		})

		e := Employee{
			Username: s.Parent().Parent().AttrOr("id", ""),
			Fullname: s.Find("strong").Text(),
			Title:    s.Find("small").Text(),
			Social:   social,
		}
		el[i] = e
	})

	return el, nil
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
