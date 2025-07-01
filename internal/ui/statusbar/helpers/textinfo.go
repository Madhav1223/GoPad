package helpers

import "strings"

type TextStats struct {
	Lines int
	Words int
	Chars int
}

func ComputeStats(text string) TextStats {
	return TextStats{
		Lines: strings.Count(text, "\n") + 1,
		Words: len(strings.Fields(text)),
		Chars: len([]rune(text)),
	}
}
