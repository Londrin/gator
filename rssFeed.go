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
	"strings"
	"time"

	"github.com/Londrin/rss-aggregator/internal/database"
	"github.com/google/uuid"
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
		return &RSSFeed{}, fmt.Errorf("new Req: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("http Req: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("readall Body: %w", err)
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

func scrapeFeeds(s *state, user database.User) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Scrape Feed: Unable to fetch next feed", err)
		return
	}

	log.Println("Found a feed")
	scrapeFeed(s.db, feed, user)
}

func scrapeFeed(db *database.Queries, feed database.Feed, user database.User) {
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

	for _, item := range feedData.Channel.Item {
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing pubDate: %v", err)
		}

		fmt.Println(item.PubDate)
		fmt.Println(pubDate)

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key value") {
				log.Println("Error creating post", err)
				return
			}
		}
	}

	fmt.Println("=====================================")
	fmt.Println()

	log.Printf("Feed %s completed, %v posts founds", feed.Name, len(feedData.Channel.Item))
}
