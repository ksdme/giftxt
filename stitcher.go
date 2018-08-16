package giftxt

import (
	"image"
	"image/gif"
	"io"
)

// MakeGif returns a GIF from a set of Paletted frames,
func MakeGif(frames []*image.Paletted, delay int) *gif.GIF {
	opGif := &gif.GIF{}

	for _, frame := range frames {
		opGif.Image = append(opGif.Image, frame)
		opGif.Delay = append(opGif.Delay, delay)
	}

	return opGif
}

// Stitch simply stitches all frames into a GIF and writes it stream.
func Stitch(frames []*image.Paletted, delay int, writer io.Writer) {
	gif.EncodeAll(writer, MakeGif(frames, delay))
}
