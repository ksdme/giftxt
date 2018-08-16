package giftxt

import (
	"io/ioutil"

	"golang.org/x/image/math/fixed"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"
)

var (
	dpi         = 72.0
	fontfile    = "../resources/WorkSans-ExtraBold.ttf"
	maxFontSize = 300.0
)

// LoadTypeFace loads a font from file and returns it.
// Font file cannot be customized.
func LoadTypeFace() *truetype.Font {
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

// MeasureStringLength measures the length of the string and the font face.
func MeasureStringLength(local *truetype.Font, size float64, text string) (fixed.Int26_6, *font.Face) {
	face := truetype.NewFace(local, &truetype.Options{
		Size: size,
		DPI:  dpi,
	})

	return font.MeasureString(face, text), &face
}

// GetAdjustedFontFace returns the adjusted font face for a given message.
func GetAdjustedFontFace(local *truetype.Font, text string, max int32) *font.Face {
	fixedMax := fixed.Int26_6(max << 6)
	fontSize := maxFontSize

	length, face := MeasureStringLength(local, fontSize, text)

	for length > fixedMax {
		fontSize -= 5
		length, face = MeasureStringLength(local, fontSize, text)
	}

	return face
}

// GetFontSizeCrossMap generates a cross reference map for the length vs. font face
func GetFontSizeCrossMap(local *truetype.Font, max int32, till int) map[int]*font.Face {
	crossMap := make(map[int]*font.Face)
	var payload string

	for length := 0; length <= till; length++ {
		crossMap[length] = GetAdjustedFontFace(local, payload, max)
		payload += "Z"
	}

	return crossMap
}

// ClampedStringLength returns the clamped length of
// the string to max cache length available.
func ClampedStringLength(text string, max int) int {
	if len(text) > max {
		return max
	}

	return len(text)
}
