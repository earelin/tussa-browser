package tussaimporter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloadAllLinesPage(t *testing.T) {
	downloadedPage := downloadAllLinesPage()

	assert.NotEmpty(t, downloadedPage)
}

func TestDownloadLinePage(t *testing.T) {
	downloadedPage := downloadLinePage(1)

	assert.NotEmpty(t, downloadedPage)
}

func TestDownloadStopsPage(t *testing.T) {
	downloadedPage := downloadStopsPage()

	assert.NotEmpty(t, downloadedPage)
}
