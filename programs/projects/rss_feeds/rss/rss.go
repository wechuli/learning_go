package rss

import (
	"errors"
	"io/ioutil"
	"net/http"
	"rss_feeds/models"
	"strings"

	"github.com/mmcdole/gofeed/rss"
	// "github.com/mmcdole/gofeed"
	// "strings"
	// "github.com/mmcdole/gofeed/rss"
)

func fetchRssFeedRaw(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {

		return "", errors.New("unable to fetch feed")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("unable to  feed")
	}
	return string(bodyBytes), nil

}

func parseRawRssString(rawRssString string) ([]models.Feed, error) {

	// data, err := fetchRssFeedRaw("http://feeds.bbci.co.uk/news/rss.xml")
	// if err != nil {
	// 	return nil, err
	// }

	fp := rss.Parser{}
	rssFeed, _ := fp.Parse(strings.NewReader(rawRssString))
	var feeds []models.Feed

	for _, item := range rssFeed.Items {

		feed := models.Feed{Title: item.Title, Description: item.Description, PubDate: item.PubDate, Link: item.Link}

		feeds = append(feeds, feed)
	}

	return feeds, nil

}
