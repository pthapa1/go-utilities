package main

import (
	"fmt"
	"regexp"

	"github.com/go-rod/rod"
)

// copied from https://www.reddit.com/r/golang/comments/12vqev1/downloadable_documentation/
func main() {
	rootUrl := "https://go.dev/"
	browser := rod.New().MustConnect()
	page := browser.MustPage(rootUrl + "doc").MustWaitLoad()

	docs := page.MustElements("section.BigCard h3 a")

	for _, doc := range docs {

		// hack the urls because i'm lazy
		// ignore any external urls
		link := *doc.MustAttribute("href")
		re := regexp.MustCompile("^http")
		if re.FindString(link) == "" {
			title := doc.MustText()
			fmt.Println(title)

			reAbs := regexp.MustCompile("^/")
			if reAbs.FindString(link) != "" {
				link = rootUrl + link
			} else {
				link = rootUrl + "doc" + link
			}
			fmt.Println(link)
			browser.MustPage(link).MustWaitLoad().MustPDF(title + ".pdf")
		}

	}

	defer browser.MustClose()
}
