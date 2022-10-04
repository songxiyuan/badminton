package third

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	Pre = "http://yuran.zhongyulian.com/activity/info/id/"
)

type ZYL struct {
}

func (z ZYL) Monitor(ids []string) (res bool, err error) {
	for _, id := range ids {
		res, err = z.check(id)
		if err != nil {
			return
		}
		if res {
			return
		}
	}
	return
}

func (ZYL) check(id string) (res bool, err error) {
	// Request the HTML page.
	resp, err := http.Get(Pre + id)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false, errors.New(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	// Find the review items
	doc.Find("body > div.container > div:nth-child(2) > div.info > div.apply-people > div.persons").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("span > div.name").Text()
		log.Println(title)
		if strings.Contains(title, os.Getenv("Name")) {
			res = true
			return
		}
	})
	return
}
