package tussaimporter

import (
	"github.com/earelin/tussa-browser/internal"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestParseAllLinesPage(t *testing.T) {
	allLinesPage, _ := ioutil.ReadFile("../../test/data/lines-page.html")

	allLines := parseAllLinesPage(string(allLinesPage))

	assert.Equal(t, expectedLines(), allLines)
}

func expectedLines() []internal.Line {
	return []internal.Line{
		{
			Id:     "26",
			Name:   "Liña 1:Cemiterio de Boisaca / Hospital Clínico",
			Routes: make([]internal.Route, 0),
		},
		{
			Id:     "3",
			Name:   "Liña 4:Romaño / Vista Alegre / As Cancelas",
			Routes: make([]internal.Route, 0),
		},
		{
			Id:     "33",
			Name:   "Liña 5:Vite / Conxo / Rocha",
			Routes: make([]internal.Route, 0),
		},
		{
			Id:     "5",
			Name:   "Liña 6:San Marcos / Os Tilos",
			Routes: make([]internal.Route, 0),
		},
	}
}
