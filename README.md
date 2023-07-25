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

imgs contains Image struct. There could be multiple images.
Primarily the scraper finds og:image for the first; if not found it will search every images available for that url and return them.

```go
type Image struct {
	Url     string `json:"url"`
	SrcType string `json:"src_type"`
}
```
