# Go Web Scraper Service

A highly efficient, scalable, and modular web scraper built with Go. This service leverages modern tools and best practices to scrape web pages and store data in a PostgreSQL database. It supports proxy-based requests for regional access and is containerized for deployment in cloud environments.

```
web_scraper_go/
├── cmd/
│   └── main.go                # Entry point of the application
├── pkg/
│   ├── api/
│   │   └── crawler.go         # Crawler logic
│   ├── config/
│   │   └── config.go          # Configuration management
│   ├── database/
│   │   └── database.go        # Database interaction
│   ├── logger/
│   │   └── logger.go          # Centralized logging
├── go.mod                     # Go module definition
├── .env                       # Environment variables
└── README.md                  # Documentation
```

---

## Features

- **High Performance**: Built with Go for lightweight and fast execution.
- **Proxy Support**: Integrates API-based proxy for country-specific scraping.
- **Database Integration**: Stores scraped content in PostgreSQL.
- **Modular Design**: Clean architecture for maintainability and scalability.
- **Dockerized Deployment**: Easily deployable on cloud platforms using Docker and Kubernetes.

---

### Prerequisites

- **Go**: Version 1.20 or later.
- **PostgreSQL**: For data storage.
- **Docker**: For containerized deployment (optional).
- **API Token**: Sign up at [Crawlbase](https://crawlbase.com/) to get your API token for proxy-based scraping.

### API Tokens
- **TOKEN_ANT2**: 4d152606c50047ff802fa196490f9d32
- **TOKEN_ANT1**: 6dbdc138f25a42ed92644a59541c77b6
