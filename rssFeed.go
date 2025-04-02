package main

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Londrin/rss-aggregator/internal/database"
)

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link,omitempty"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("New Req: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Http Req: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Readall Body: %w", err)
	}

	feed := RSSFeed{}
	decoder := xml.NewDecoder(bytes.NewReader(body))
	err = decoder.Decode(&feed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("XML Unpack: %w", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i] = item
	}

	return &feed, nil
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Scrape Feed: Unable to fetch next feed", err)
		return
	}

	log.Println("Found a feed")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Scrape Feed: Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Scrape Feed: Unable to fetch feed %s: %v", feed.Name, err)
		return
	}

	fmt.Printf("Feed Name: %s\n", feedData.Channel.Title)
	fmt.Printf("Feed Link: %s\n", feedData.Channel.Link)
	fmt.Println()
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Title: %s\n", item.Title)
	}

	fmt.Println("=====================================")
	fmt.Println()

	log.Printf("Feed %s completed, %v posts founds", feed.Name, len(feedData.Channel.Item))
}
