package third

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	Root          = "http://yuran.zhongyulian.com"
	ActivityIndex = "/Activity/index/typeid/0/day/"
)

type ZYL struct {
}

func (z ZYL) Monitor(days []string) (hit bool, url string) {
	for _, day := range days {
		hit, url = z.check(day)
		if hit {
			return hit, Root + url
		}
	}
	return
}

func (ZYL) check(day string) (hit bool, url string) {
	// Request the HTML page.
	resp, err := http.Get(Root + ActivityIndex + day)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	// Find the review items
	activities := doc.Find("body > div.container > div:nth-child(2) > div.right > div.activities-list")
	activities.Find("div.list").Each(func(i int, activitySelection *goquery.Selection) {
		activitySelection.Find("span.apply-number > div.apply-people").Find("img").Each(func(j int, nameSelection *goquery.Selection) {
			name, _ := nameSelection.Attr("title")
			if strings.Contains(name, os.Getenv("Name")) {
				url, _ = activitySelection.Find("a").Attr("href")
				hit = true
				return
			}
		})
	})

	return
}

/*


	doc.Find("body > div.container > div:nth-child(2) > div.right > div.activities-list").Each(func(i int, s1 *goquery.Selection) {
		// For each item found, get the title
		s1.Find("span.apply-number > div.apply-people").Find("img").Each(func(j int, s2 *goquery.Selection) {
			name, _ := s2.Attr("title")
			if strings.Contains(name, os.Getenv("Name")) {
				url, _ = s1.Find("a").Attr("href")
				fmt.Println(url)
				hit = true
				return
			}
		})
	})
*/
