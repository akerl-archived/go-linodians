package api

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Employee describes a human at Linode
type Employee struct {
	Username string
	Fullname string
	Title    string
	Social   map[string]string
}

func employeeFromDiv(s *goquery.Selection) Employee {
	social := make(map[string]string)
	s.Find("a.employee-link").Each(func(_ int, ss *goquery.Selection) {
		class := ss.AttrOr("class", "")
		parts := strings.Split(class, "-")
		site := parts[len(parts)-1]
		link := ss.AttrOr("href", "")
		social[site] = link
	})
	return Employee{
		Username: s.Parent().Parent().AttrOr("id", ""),
		Fullname: s.Find("strong").Text(),
		Title:    s.Find("small").Text(),
		Social:   social,
	}
}

// Photo downloads the photo for an employee
func (e Employee) Photo() ([]byte, error) {
	if e.Username == "" {
		return []byte{}, fmt.Errorf("Username not set")
	}
	url := fmt.Sprintf(photoURLFormat, e.Username)
	return download(url)
}
