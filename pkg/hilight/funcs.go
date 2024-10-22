package hilight

import "strings"

func HighlightKeyword(keyword string, str string, color string) string {
	return strings.ReplaceAll(str, keyword, color+keyword+Reset)
}
