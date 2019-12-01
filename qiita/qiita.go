package qiita

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/tockn/vs-dena-advent/model"
)

const baseURL = "https://qiita.com/advent-calendar/%d/%s"

func getAdventDoc(year int64, title string) (*goquery.Document, error) {
	URL := fmt.Sprintf(baseURL, year, title)
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(res.Body)
}

func GetAllLikes(year int64, title string) (int64, error) {
	doc, err := getAdventDoc(year, title)
	if err != nil {
		return 0, err
	}
	selection := doc.Find("div.adventCalendarJumbotron_stats[title=Likes]")
	likesStr := selection.Text()
	likesStr = strings.TrimSpace(likesStr)
	return strconv.ParseInt(likesStr, 10, 64)
}

func GetArticles(year int64, title string) ([]model.Article, error) {
	doc, err := getAdventDoc(year, title)
	if err != nil {
		return nil, err
	}
	authorSel := doc.Find("div.adventCalendarCalendar_author")
	titleSel := doc.Find("div.adventCalendarCalendar_comment")
	as := make([]model.Article, 25)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		authorSel.Each(func(i int, selection *goquery.Selection) {
			as[i].Author = strings.TrimSpace(selection.Text())
		})
		wg.Done()
	}()
	go func() {
		titleSel.Each(func(i int, selection *goquery.Selection) {
			as[i].Title = strings.TrimSpace(selection.Text())
			as[i].URL, _ = selection.Find("a").Attr("href")
			if as[i].URL == "" {
				return
			}
			split := strings.Split(as[i].URL, "/")
			articleID := split[len(split)-1]
			as[i].Likes, _ = GetLikesByArticleID(articleID)
		})
		wg.Done()
	}()
	wg.Wait()
	return as, nil
}

func GetLikesByArticleID(id string) (int64, error) {
	res, err := http.Get(fmt.Sprintf("https://qiita.com/api/v2/items/%s/likes", id))
	if err != nil {
		return 0, err
	}
	totalCount := res.Header["Total-Count"]
	if len(totalCount) == 0 {
		return 0, nil
	}
	return strconv.ParseInt(totalCount[0], 10, 64)
}
