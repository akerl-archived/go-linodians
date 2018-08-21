package api

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

// Company describes a map of Employees by username
type Company map[string]Employee

func companyFromDivFunc(c Company) func(int, *goquery.Selection) {
	return func(_ int, s *goquery.Selection) {
		e := employeeFromDiv(s)
		c[e.Username] = e
	}
}

// Load employees from site
func Load() (Company, error) {
	page, err := download(dataURL)
	if err != nil {
		return Company{}, err
	}

	reader := bytes.NewReader(page)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return Company{}, err
	}

	divs := doc.Find("div.employee-display")
	c := make(map[string]Employee, len(divs.Nodes))
	divs.Each(companyFromDivFunc(c))
	return c, nil
}
