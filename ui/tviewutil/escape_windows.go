package tviewutil

import (
	tview "github.com/gord-project/gview"
	"github.com/mattn/go-runewidth"
	"github.com/yellowsink/gord/config"
)

// Escape delegates to tview escape, optionally doing additional escaping.
func Escape(text string) string {
	if config.DisableUTF8 {
		runes := []rune(text)
		for index, r := range runes {
			if r > 65536 || runewidth.RuneWidth(r) > 1 {
				runes[index] = '?'
			}
		}

		return tview.Escape(string(runes))
	}

	return tview.Escape(text)
}
