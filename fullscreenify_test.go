package fullscreenify

import (
	"fmt"
	"testing"
)

func TestYoutube(t *testing.T) {
	tests := [][]string{
		[]string{"https://www.youtube.com/watch?v=ys0HyevZpQg&t=1849s", "https://www.youtube.com/embed/ys0HyevZpQg?rel=0&autoplay=1"},
		[]string{"https://www.youtube.com/watch?v=ys0HyevZpQg", "https://www.youtube.com/embed/ys0HyevZpQg?rel=0&autoplay=1"},
		[]string{"https://youtu.be/ys0HyevZpQg", "https://www.youtube.com/embed/ys0HyevZpQg?rel=0&autoplay=1"},
		[]string{"https://youtu.be/IyzdDu7ZHz8?autoplay=1", "https://www.youtube.com/embed/IyzdDu7ZHz8?rel=0&autoplay=1"},
	}

	for i := range tests {
		r := Fullscreenify(tests[i][0])
		if r == tests[i][1] {
			fmt.Printf("t%d passed :)\n", i)
		} else {
			t.Fatalf("t%d failed\n", i)
		}
	}
}
