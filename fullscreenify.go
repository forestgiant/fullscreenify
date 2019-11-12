package fullscreenify

import (
	"fmt"
	"net/url"
	"strings"
)

type Parser func(string) string

var parsers map[string]Parser

func youtube(urlString string) string {
	template := "https://www.youtube.com/embed/%s?rel=0&autoplay=1"
	startIdx := strings.Index(urlString, "v=") + 2
	endIdx := startIdx + 11
	return fmt.Sprintf(template, urlString[startIdx:endIdx])
}

func youtuDotBe(urlString string) string {
	template := "https://www.youtube.com/embed/%s?rel=0&autoplay=1"
	uURL, err := url.Parse(urlString)
	if err != nil {
		return ""
	}
	return fmt.Sprintf(template, uURL.Path[1:])
}

func init() {
	parsers = make(map[string]Parser)
	parsers["www.youtube.com"] = youtube
	parsers["youtu.be"] = youtuDotBe
}

func Fullscreenify(urlStr string) string {
	tURL, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}
	//TODO: make sure hostname is in the map
	//TODO: if not in the map or error return the original
	return parsers[tURL.Hostname()](urlStr)
}
