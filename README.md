# 👍 go-meta-scraper

Simple go-meta-scraper in Go with surf and goquery.
scrapes thumbnail images from given site url, with dynamic crawling(using headless browser).

## More

- [surf](https://github.com/headzoo/surf)
- [goquery](https://github.com/PuerkitoBio/goquery)

## Usage

```go
package main

import "github.com/chaewonkong/thumbnailscraper"

func main() {
	url := [YOUR_URL]

	// when
	ts := thumbnailscraper.New(url)
	imgs, err := ts.GetThumbnailImages()
}
```
