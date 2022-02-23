package main

import (
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/nathan-fiscaletti/consolesize-go"
)

var termWidth = 0

func calculateWidth() {
	if os.Getenv("FZF_PREVIEW_COLUMNS") != "" {
		i, err := strconv.ParseInt(os.Getenv("FZF_PREVIEW_COLUMNS"), 10, 64)
		panicIfErr(err)
		termWidth = int(i)
	} else {
		width, _ := consolesize.GetConsoleSize()
		termWidth = width
	}
}

func sperator(left string, right string) string {
	leftW := termWidth
	leftW -= 2
	str := left
	for leftW >= 1 {
		str += "─"
		leftW--
	}
	return str + right + "\n"
}

func centerThings(text string, prefix string, suffix string) string {
	left := (termWidth - utf8.RuneCountInString(text) - utf8.RuneCountInString(prefix) - utf8.RuneCountInString(suffix))/2
	space := ""
	for left >= 1 {
		space += " "
		left--
	}
	extra := ""
	if (utf8.RuneCountInString(text) + utf8.RuneCountInString(prefix) + utf8.RuneCountInString(suffix) + len(space)*2) != termWidth {
		extra += " "
	}
	
	return prefix + space + text + extra + space + suffix + "\n"
}

func centerText(text string) string {
	return centerThings(text, "│", "│")
}