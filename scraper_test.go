package thumbnailscraper_test

import (
	"testing"

	thumbnailscraper "github.com/chaewonkong/thumbnail-scraper"
	"github.com/stretchr/testify/assert"
)

func TestGetThumbnailImagesIntegrationTest(t *testing.T) {
	// given
	url := "https://leonkong.cc/"
	expected := thumbnailscraper.Image{Url: "/tinamn.jpg", SrcType: "og:image"}

	// when
	ts := thumbnailscraper.New(url)
	imgs, err := ts.GetThumbnailImages()

	// then
	assert.NoError(t, err)
	assert.Contains(t, imgs, expected)
}
