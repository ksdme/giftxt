package main

import (
	"image/gif"
	"log"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/ksdme/giftxt"
)

const (
	frameSize    = 300 // Requires change at frame.go
	textPortSize = int32(frameSize * 0.8)
	cacheTill    = 30
	delay        = 50
)

// Preload all the referables
var typeface = giftxt.LoadTypeFace()
var fontSizeCrossMap = giftxt.GetFontSizeCrossMap(
	typeface, textPortSize, cacheTill)

// Generates a new GifText from a given text
func generateGifText(text string) *gif.GIF {
	message := giftxt.NewMessage(text)
	length := giftxt.ClampedStringLength(message.Longest, cacheTill)

	face, _ := fontSizeCrossMap[length]
	slides := giftxt.RenderMessage(face, message.Words)

	return giftxt.MakeGif(slides, delay)
}

func main() {
	// Serving Port
	port := os.Getenv("PORT")

	if len(port) == 0 {
		log.Fatal("PORT environment variable missing")
	}

	app := iris.Default()

	// Method:   GET;
	// /generate takes a text string and returns a GIF.
	app.Get("/generate", func(ctx context.Context) {
		text := ctx.URLParamDefault("text", "hello world")
		giftxt := generateGifText(text)

		gif.EncodeAll(ctx.ResponseWriter(), giftxt)
		ctx.ContentType("image/gif")
	})

	app.Run(iris.Addr(":" + port))
}
