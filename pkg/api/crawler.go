package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type CrawlerAPI struct {
	Token string
}

var COUNTRY_CODES = map[string]string{
	"All countries":        "",
	"USA":                  "US",
	"United Arab Emirates": "AE",
	"Brazil":               "BR",
	"Canada":               "CA",
	"China":                "CN",
	"Czech Republic":       "CZ",
	"Germany":              "DE",
	"Spain":                "ES",
	"France":               "FR",
	"United Kingdom":       "GB",
	"Hong Kong":            "HK",
	"India":                "IN",
	"Italy":                "IT",
	"Israel":               "IL",
	"Japan":                "JP",
	"Netherlands":          "NL",
	"Poland":               "PL",
	"Russia":               "RU",
	"Saudi Arabia":         "SA",
	"Singapore":            "SG",
	"South Korea":          "KR",
	"Indonesia":            "ID",
	"Vietnam":              "VN",
}

// Initialize CrawlerAPI with a token
func NewCrawlerAPI(token string) *CrawlerAPI {
	return &CrawlerAPI{Token: token}
}

// Function to save the response to a file
func (c *CrawlerAPI) writeToFile(inputFile string, data string) error {
	err := ioutil.WriteFile(inputFile, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}
	log.Printf("Successfully saved response to %s", inputFile)
	return nil
}

// Fetch webpage using Crawlbase
func (c *CrawlerAPI) FetchAndSaveWebpage(urlToFetch, inputFile, proxyCountry string) error {
	encodedURL := url.QueryEscape(urlToFetch)
	requestURL := fmt.Sprintf("https://api.crawlbase.com/?token=%s&url=%s", c.Token, encodedURL)

	// Check for proxy country and append to the URL
	if proxyCountryCode, exists := COUNTRY_CODES[proxyCountry]; exists && proxyCountryCode != "" {
		requestURL += fmt.Sprintf("&proxy_country=%s", proxyCountryCode)
	}

	// Make the HTTP request
	resp, err := http.Get(requestURL)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Save the response to a file
	return c.writeToFile(inputFile, string(body))
}

// Fetch webpage using ScrapingAnt
func (c *CrawlerAPI) FetchWebpageAnt(urlToFetch, inputFile, proxyCountry string) error {
	encodedURL := url.QueryEscape(urlToFetch)
	requestURL := fmt.Sprintf("/v2/general?url=%s&x-api-key=%s&return_page_source=true", encodedURL, c.Token)

	// Check for proxy country and append to the URL
	if proxyCountryCode, exists := COUNTRY_CODES[proxyCountry]; exists && proxyCountryCode != "" {
		requestURL += fmt.Sprintf("&proxy_country=%s", proxyCountryCode)
	}

	// Make the HTTP request to ScrapingAnt
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.scrapingant.com"+requestURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Save the response to a file
	return c.writeToFile(inputFile, string(body))
}
