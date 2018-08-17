package giftxt

import (
	"io/ioutil"

	"golang.org/x/image/math/fixed"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"
)

var (
	dpi         = 72.0
	maxFontSize = 300.0
)

// LoadTypeFace loads a font from file and returns it.
// Font file cannot be customized.
func LoadTypeFace(fontfile string) *truetype.Font {
	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		panic("cannot read font file")
	}

	fontParsed, err := truetype.Parse(fontBytes)
	if err != nil {
		panic("parsing font file failed")
	}

	return fontParsed
}

// GetFaceFromFontAndSize returns a face from a font and font size.
func GetFaceFromFontAndSize(local *truetype.Font, fontSize float64) *font.Face {
	face := truetype.NewFace(local, &truetype.Options{
		DPI:  dpi,
		Size: fontSize,
	})

	return &face
}

// MeasureStringLength measures the length of the string and the font face.
func MeasureStringLength(local *truetype.Font, size float64, text string) (fixed.Int26_6, *font.Face) {
	face := GetFaceFromFontAndSize(local, size)
	return font.MeasureString(*face, text), face
}

// GetAdjustedFontSize returns the adjusted font size for a given message.
func GetAdjustedFontSize(local *truetype.Font, text string, max int32) float64 {
	fixedMax := fixed.Int26_6(max << 6)
	fontSize := maxFontSize

	length, _ := MeasureStringLength(local, fontSize, text)

	for length > fixedMax {
		fontSize -= 4
		length, _ = MeasureStringLength(local, fontSize, text)
	}

	return fontSize
}
