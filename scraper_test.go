package surfio_test

import (
	"testing"

	"github.com/chaewonkong/surfio"
	"github.com/chaewonkong/surfio/property"
	"github.com/stretchr/testify/assert"
)

func TestGetThumbnailImagesIntegrationTest(t *testing.T) {
	// given
	url := "https://velog.io/@superlipbalm/how-i-use-adapter-pattern-in-reactjs"

	expected := "/tinamn.jpg"

	// when
	ts := surfio.New(url)
	imgs, err := ts.GetMetaContentByProp(property.OGImage)

	titles, err := ts.GetMetaContentByProp(property.OGTitle)

	_ = titles

	assert.NoError(t, err)
	// then
	assert.Contains(t, imgs, expected)
}

func TestGetMetaContentByProp(t *testing.T) {
	url := "https://leonkong.cc"
	t.Run("Finds og:image successfully", func(t *testing.T) {
		// given
		expected := "/tinamn.jpg"
		// when
		ts := surfio.New(url)
		imgs, err := ts.GetMetaContentByProp(property.OGImage)

		// then
		assert.NoError(t, err)

		// then
		assert.Contains(t, imgs, expected)
	})

	t.Run("Finds title successfully", func(t *testing.T) {
		// given
		expected := "LEON KONG :: Leon Kong's Blog"
		// when
		ts := surfio.New(url)
		titles, err := ts.GetMetaContentByProp(property.OGTitle)

		// then
		assert.NoError(t, err)

		// then
		assert.Contains(t, titles, expected)
	})

	t.Run("Finds description successfully", func(t *testing.T) {
		// given
		expected := "A Blog For tech, IT, Computer Engineering, Backend and Web"
		// when
		ts := surfio.New(url)
		desc, err := ts.GetMetaContentByProp(property.OGDescription)

		// then
		assert.NoError(t, err)

		// then
		assert.Contains(t, desc, expected)
	})
}
