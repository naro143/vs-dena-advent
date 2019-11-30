package qiita

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

const baseURL = "https://qiita.com/advent-calendar/%d/%s"

func GetLikes(year int64, title string) (int64, error) {
	URL := fmt.Sprintf(baseURL, year, title)
	res, err := http.Get(URL)
	if err != nil {
		return 0, err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return 0, err
	}
	selection := doc.Find("div.adventCalendarJumbotron_stats[title=Likes]")
	likesStr := selection.Text()
	likesStr = strings.TrimSpace(likesStr)
	return strconv.ParseInt(likesStr, 10, 64)
}
