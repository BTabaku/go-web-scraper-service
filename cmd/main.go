// cmd/main.go
package main

import (
	"fmt"
	"log"
	"web_scraper_go/pkg/api"
	"web_scraper_go/pkg/config"
	"web_scraper_go/pkg/database"
	"web_scraper_go/pkg/logger"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db, err := database.NewDatabase(cfg.DBConnectionString)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Initialize crawler
	crawler := api.NewCrawler(cfg.APIToken, db)

	// Accept user input
	var url, proxyCountry string
	log.Println("Enter the URL to scrape:")
	fmt.Scanln(&url)
	log.Println("Enter proxy country (optional):")
	fmt.Scanln(&proxyCountry)

	// Fetch and save webpage
	err = crawler.FetchAndSaveWebpage(url, proxyCountry)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
