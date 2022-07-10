package tussaimporter

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/earelin/tussa-browser/internal"
	"html"
	"regexp"
	"strings"
)

var findId, _ = regexp.Compile("linas.php\\?id=([0-9]*)")

func parseAllLinesPage(page string) []internal.Line {
	pageDocument, _ := goquery.NewDocumentFromReader(strings.NewReader(page))

	return extractLinesFromDocument(pageDocument)
}

func extractLinesFromDocument(doc *goquery.Document) []internal.Line {
	lines := make([]internal.Line, 0)

	doc.Find(".linasbus li").Each(func(i int, s *goquery.Selection) {
		line := internal.Line{
			Id:     extractId(s),
			Name:   extractName(s),
			Routes: make([]internal.Route, 0),
		}
		lines = append(lines, line)
	})

	return lines
}

func extractId(s *goquery.Selection) string {
	idText, _ := s.Find("a").Attr("href")
	matches := findId.FindStringSubmatch(idText)
	return matches[1]
}

func extractName(s *goquery.Selection) string {
	trimmedText := strings.TrimSpace(s.Text())
	return html.UnescapeString(trimmedText)
}
