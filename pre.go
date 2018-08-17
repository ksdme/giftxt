package giftxt

import (
	"encoding/json"
	"io"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// FontFaceCache is struct that holds a
// precalculated set of font faces.
type FontFaceCache struct {
	CacheMap   map[int]float64
	Till       int
	CanvasSize int32
}

// ClampedInt returns a clamped integer where if value > max
// it returns max or if val < min then it returns min.
func ClampedInt(value, min, max int) int {
	if value > max {
		return value
	}

	if value < min {
		return value
	}

	return value
}

// ClampedStringLength returns the clamped length of
// the string to max cache length available.
func ClampedStringLength(text string, max int) int {
	return ClampedInt(len(text), 0, max)
}

// GetFontSizeCrossMap generates a cross reference map for the length vs. font face
func GetFontSizeCrossMap(local *truetype.Font, size int32, till int) *FontFaceCache {
	crossMap := make(map[int]float64)
	var payload string

	for length := 0; length <= till; length++ {
		crossMap[length] = GetAdjustedFontSize(local, payload, size)
		payload += "Z"
	}

	return &FontFaceCache{
		CacheMap:   crossMap,
		Till:       till,
		CanvasSize: size,
	}
}

// GetFontSize returns a font size from font size.
func (context *FontFaceCache) GetFontSize(length int) float64 {
	length = ClampedInt(length, 1, context.Till)
	fontSize, ok := context.CacheMap[length]

	if !ok {
		fontSize, _ = context.CacheMap[1]
		log.Println("failed accessing font face", ok)
	}

	return fontSize
}

// GetFace returns a font face from font size.
func (context *FontFaceCache) GetFace(local *truetype.Font, length int) *font.Face {
	return GetFaceFromFontAndSize(local, context.GetFontSize(length))
}

// Export serializes the current object in gob encoding
func (context *FontFaceCache) Export(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(context)
}

// LoadFontFaceCache loads a FontFaceCache and returns it
func LoadFontFaceCache(reader io.Reader) (*FontFaceCache, error) {
	var loaded FontFaceCache

	decoder := json.NewDecoder(reader)
	ok := decoder.Decode(&loaded)

	return &loaded, ok
}
