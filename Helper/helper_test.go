//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package Helper

import (
	"image"
	"image/color"
	"os"
	"testing"
)

func TestImageWriter(t *testing.T) {
	width := 640
	height := 480
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{width, height}})
	//starts from left side of the image for each pixel to the right side of the image
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if (x > 200 && x < 440) && (y > 200 && y < 280) {
				m.SetRGBA(x, y, color.RGBA{0, 0, 255, 255})
			} else {
				m.SetRGBA(x, y, color.RGBA{0, 0, 0, 255})
			}
		}
	}

	filename := "testpicture.png"
	h := new(Helper)
	h.ImageWriter(filename, m)
	if _, err := os.Stat(filename); err != nil {
		t.Errorf("TestImageWriter: %s", err)
	}

	filename = "T:\\TTTTTT"
	h.ImageWriter(filename, m)
}

func TestRound(t *testing.T) {
	if Round(1.0, 0) != 1.0 {
		t.Errorf("TestRound: %f")
	}

	if Round(0.5, 0) != 1.0 {
		t.Errorf("TestRound: %f %f")
	}

	if Round(0.49, 0) != 0.0 {
		t.Errorf("TestRound: %f")
	}
}
