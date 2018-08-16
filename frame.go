package giftxt

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	dimension   = 300
	colorFg     = color.White
	colorBg     = color.Black
	alternating = false
)

// Used to cache the value generated.
var textSrc = GetColoredImage(colorFg)

// NewBlankImage returns a blank new Image
func NewBlankImage() *image.Paletted {
	return image.NewPaletted(image.Rect(0, 0, dimension, dimension), palette.Plan9)
}

// GetColoredImage returns base image to use while drawing.
func GetColoredImage(paint color.Color) *image.Paletted {
	frame := NewBlankImage()
	draw.Draw(frame, frame.Bounds(), image.NewUniform(paint), image.ZP, draw.Src)

	return frame
}

// RenderSingleWord renders the word onto an image with colorBg
func RenderSingleWord(face *font.Face, word string) *image.Paletted {
	dst := GetColoredImage(colorBg)

	bounds, _ := font.BoundString(*face, word)
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	left := (fixed.Int26_6(dimension<<6) - width) / 2
	top := (fixed.Int26_6(dimension<<6) + height) / 2

	drawer := &font.Drawer{
		Src:  textSrc,
		Dst:  dst,
		Face: *face,
		Dot: fixed.Point26_6{
			X: left,
			Y: top,
		},
	}

	drawer.DrawString(word)
	return dst
}

// RenderMessage renders the entire message and returns frames
func RenderMessage(face *font.Face, words []string) []*image.Paletted {
	var frames []*image.Paletted

	for _, word := range words {
		frames = append(frames, RenderSingleWord(face, word))
	}

	return frames
}
