# thumbnail-scraper

Simple thumbnail-scraper in Go with surf and goquery.

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
