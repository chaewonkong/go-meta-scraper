# 🌊 surfio

Simple image scraper in Go with surf and goquery.
Scrapes thumbnail images from given site url, with dynamic crawling(using headless browser).

## More

- [surf](https://github.com/headzoo/surf)
- [goquery](https://github.com/PuerkitoBio/goquery)

## Usage

```go
package main

import "github.com/chaewonkong/surfio"

func main() {
	url := [YOUR_URL]

	// when
	su := surfio.New(url)
	imgs, err := su.GetThumbnailImages()

	if err != nil {
		// handle error
	}

	// if not error, use image
	imgUrl := imgs[0].Url
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

## Todo

- [ ] Crawling Site meta (title, description, keyword...)
