package surfio_test

import (
	"testing"

	"github.com/chaewonkong/surfio"
	"github.com/stretchr/testify/assert"
)

func TestGetThumbnailImagesIntegrationTest(t *testing.T) {
	// given
	url := "https://leonkong.cc/"

	expected := surfio.Image{Url: "/tinamn.jpg", SrcType: "og:image"}

	// when
	ts := surfio.New(url)
	imgs, err := ts.GetThumbnailImages()

	// then
	assert.NoError(t, err)
	assert.Contains(t, imgs, expected)
}
