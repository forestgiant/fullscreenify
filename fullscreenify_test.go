package fullscreenify

import (
	"flag"
	"fmt"
	"testing"
)

var verbose = flag.Bool("verbose", false, "Add verbosity")

func TestYoutube(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"Test youtube v= with parameter",
			"https://www.youtube.com/watch?v=ys0HyevZpQg&t=1849s",
			"https://www.youtube.com/embed/ys0HyevZpQg?rel=0&autoplay=1",
		},
		{
			"Test youtube v=",
			"https://www.youtube.com/watch?v=ys0HyevZpQg",
			"https://www.youtube.com/embed/ys0HyevZpQg?rel=0&autoplay=1",
		},
		{
			"Test youtu.be",
			"https://youtu.be/ys0HyevZpQg",
			"https://www.youtube.com/embed/ys0HyevZpQg?rel=0&autoplay=1",
		},
		{
			"Test yout.be w/ parameter",
			"https://youtu.be/IyzdDu7ZHz8?autoplay=1",
			"https://www.youtube.com/embed/IyzdDu7ZHz8?rel=0&autoplay=1",
		},
	}

	for _, test := range tests {
		r := ToFullscreen(test.input)
		if r == test.expected {
			if *verbose {
				fmt.Printf("%s passed\n", test.name)
			}
		} else {
			t.Fatalf("%s failed\n", test.name)
		}
	}
}

func TestGiphy(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"Test giphy",
			"https://giphy.com/gifs/positive-vibes-l41m0MZzfp7yEND1K",
			"https://giphy.com/gifs/positive-vibes-l41m0MZzfp7yEND1K/fullscreen",
		},
		{
			"Test giphy fullscreen",
			"https://giphy.com/gifs/positive-vibes-l41m0MZzfp7yEND1K/fullscreen",
			"https://giphy.com/gifs/positive-vibes-l41m0MZzfp7yEND1K/fullscreen",
		},
		{
			"Test giphy tile",
			"https://giphy.com/gifs/positive-vibes-l41m0MZzfp7yEND1K/tile",
			"https://giphy.com/gifs/positive-vibes-l41m0MZzfp7yEND1K/tile",
		},
		{
			"Test giphy media",
			"https://media.giphy.com/media/l41m0MZzfp7yEND1K/giphy.gif",
			"https://giphy.com/gifs/l41m0MZzfp7yEND1K/fullscreen",
		},
		{
			"Test giphy media social",
			"https://media.giphy.com/media/t7Rl9ag98MDRXsre46/giphy-downsized-large.gif",
			"https://giphy.com/gifs/t7Rl9ag98MDRXsre46/fullscreen",
		},
		{
			"Test giphy media mp4",
			"https://media.giphy.com/media/t7Rl9ag98MDRXsre46/giphy.mp4",
			"https://giphy.com/gifs/t7Rl9ag98MDRXsre46/fullscreen",
		},
		{
			"Test giphy media small",
			"https://media.giphy.com/media/t7Rl9ag98MDRXsre46/200w_d.gif",
			"https://giphy.com/gifs/t7Rl9ag98MDRXsre46/fullscreen",
		},
		{
			"Test giphy html5",
			"https://giphy.com/gifs/l41m0MZzfp7yEND1K/html5",
			"https://giphy.com/gifs/l41m0MZzfp7yEND1K/fullscreen",
		},
		{
			"Test gph.is",
			"http://gph.is/1CiFXmZ",
			"http://gph.is/1CiFXmZ",
		},
	}

	for _, test := range tests {
		r := ToFullscreen(test.input)
		if r == test.expected {
			if *verbose {
				fmt.Printf("%s passed\n", test.name)
			}
		} else {
			t.Fatalf("%s failed\n", test.name)
		}
	}
}

func TestVimeo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"Test vimeo",
			"https://vimeo.com/76979871",
			"https://player.vimeo.com/video/76979871?autoplay=1&loop=1&autopause=0",
		},
		{
			"Test vimeo fullscreen",
			"https://player.vimeo.com/video/76979871?autoplay=1&loop=1&autopause=0",
			"https://player.vimeo.com/video/76979871?autoplay=1&loop=1&autopause=0",
		},
	}

	for _, test := range tests {
		r := ToFullscreen(test.input)
		if r == test.expected {
			if *verbose {
				fmt.Printf("%s passed\n", test.name)
			}
		} else {
			t.Fatalf("%s failed\n", test.name)
		}
	}
}

func TestTwitch(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"Test twitch",
			"https://www.twitch.tv/ourlifewithewe",
			"https://player.twitch.tv/?channel=ourlifewithewe",
		},
		{
			"Test vimeo fullscreen",
			"https://player.twitch.tv/?channel=ourlifewithewe",
			"https://player.twitch.tv/?channel=ourlifewithewe",
		},
	}

	for _, test := range tests {
		r := ToFullscreen(test.input)
		if r == test.expected {
			if *verbose {
				fmt.Printf("%s passed\n", test.name)
			}
		} else {
			t.Fatalf("%s failed\n", test.name)
		}
	}
}
