// pkg/api/crawler.go
package api

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"web_scraper_go/pkg/database"
)

type Crawler struct {
	APIToken string
	DB       *database.Database
}

func NewCrawler(apiToken string, db *database.Database) *Crawler {
	return &Crawler{APIToken: apiToken, DB: db}
}

func (c *Crawler) FetchAndSaveWebpage(pageURL, proxyCountry string) error {
	if pageURL == "" {
		return errors.New("page URL is required")
	}

	// Construct API request
	encodedURL := url.QueryEscape(pageURL)
	apiURL := fmt.Sprintf("https://api.crawlbase.com/?token=%s&url=%s", c.APIToken, encodedURL)
	if proxyCountry != "" {
		apiURL += "&country=" + proxyCountry
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("failed to fetch page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = c.DB.SavePage(pageURL, string(body))
	if err != nil {
		return fmt.Errorf("failed to save page: %w", err)
	}

	log.Printf("Page saved successfully: %s", pageURL)
	return nil
}
