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

func giphy(urlString string) string {
	template := "https://giphy.com%s/fullscreen"
	uURL, err := url.Parse(urlString)
	if err != nil {
		return ""
	}
	pathParts := strings.Split(uURL.Path, "/")
	lastPart := pathParts[len(pathParts)-1]
	if lastPart == "fullscreen" || lastPart == "tile" {
		return urlString
	}
	if lastPart == "html5" {
		template := "https://giphy.com/gifs/%s/fullscreen"
		penultimatePart := pathParts[len(pathParts)-2]
		return fmt.Sprintf(template, penultimatePart)
	}

	return fmt.Sprintf(template, uURL.Path)
}

func giphyMedia(urlString string) string {
	template := "https://giphy.com/gifs/%s/fullscreen"
	uURL, err := url.Parse(urlString)
	if err != nil {
		return ""
	}
	pathParts := strings.Split(uURL.Path, "/")
	penultimatePart := pathParts[len(pathParts)-2]

	return fmt.Sprintf(template, penultimatePart)
}

func vimeo(urlString string) string {
	template := "https://player.vimeo.com/video/%s?autoplay=1&loop=1&autopause=0"
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
	parsers["giphy.com"] = giphy
	parsers["media.giphy.com"] = giphyMedia
	parsers["vimeo.com"] = vimeo
}

//Fullscreenify given a url return the fullscreened version of that URL
func Fullscreenify(urlStr string) string {
	uURL, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}

	if parser, ok := parsers[uURL.Hostname()]; ok {
		return parser(urlStr)
	} else {
		return urlStr
	}
}
