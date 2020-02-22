package models

import (
	"database/sql"
)

// Feed struct
type Feed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PubDate     string `json:"pubdate"`
	Link        string `json:"link"`
}

// fetch all rss feeds given a search string
func searchRssFeeds(db *sql.DB, searchString string) ([]Feed, error) {
	rows, err := db.Query("SELECT title,description,pubdate,link FROM feed")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	feeds := []Feed{}

	for rows.Next() {
		var f Feed
		if err := rows.Scan(&f.Title, &f.Description, &f.PubDate, &f.Link); err != nil {
			return nil, err
		}
		feeds = append(feeds, f)
	}
	return feeds, nil
}

// store feeds in the db

func storeRssFeeds(db *sql.DB, feeds []Feed) error {
	for _, item := range feeds {
		// skip if a row already contains the link - duplicated rss feed
		_, err := db.Query("INSERT INTO feeds(title,description,pubdate,link) VALUES($1,$2,$3,$4)", item.Title, item.Description, item.PubDate, item.Link)
		if err != nil {
			return err
		}

	}
	return nil
}
