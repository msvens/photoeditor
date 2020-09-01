package editor

import (
	"image"
	"math"
)

type Editor struct {
	thumb     image.Rectangle
	maxWidth  int
	square    image.Rectangle
	landscape image.Rectangle
	portrait  image.Rectangle
	quality   int
}

func NewEditor(maxWidth, thumbX, thumbY int) *Editor {
	//landscape 1.91:1
	y := int(math.Round(float64(maxWidth) / 1.91))
	landscape := image.Rect(0, 0, maxWidth, y)

	//portrait: 4:5
	y = int(math.Round(float64(maxWidth) / 0.8))
	portrait := image.Rect(0, 0, maxWidth, y)

	return &Editor{
		quality:   90,
		maxWidth:  maxWidth,
		thumb:     image.Rect(0, 0, thumbX, thumbY),
		landscape: landscape,
		portrait:  portrait,
		square:    image.Rect(0, 0, maxWidth, maxWidth),
	}
}

func InstaEditor(thumbX, thumbY int) *Editor {
	return &Editor{
		quality:   90,
		thumb:     image.Rect(0, 0, thumbX, thumbY),
		maxWidth:  1200,
		square:    image.Rect(0, 0, 1200, 1200),
		portrait:  image.Rect(0, 0, 1080, 1350),
		landscape: image.Rect(0, 0, 1200, 628),
	}
}

type option func(*Editor)

func (e *Editor) Option(options ...option) {
	for _, option := range options {
		option(e)
	}
}

func Quality(quality int) option {
	return func(e *Editor) {
		e.quality = quality
	}
}

func Portrait(width, height int) option {
	return func(e *Editor) {
		e.portrait = image.Rect(0, 0, width, height)
	}
}

func Landscape(width, height int) option {
	return func(e *Editor) {
		e.landscape = image.Rect(0, 0, width, height)
	}
}

func Square(width, height int) option {
	return func(e *Editor) {
		e.square = image.Rect(0, 0, width, height)
	}
}
func ThumbSize(width, height int) option {
	return func(e *Editor) {
		e.thumb = image.Rect(0, 0, width, height)
	}
}

func FullSize(size int) option {
	return func(e *Editor) {
		e.maxWidth = size
	}
}
