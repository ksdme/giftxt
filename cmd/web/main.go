package main

import (
	"flag"
	"image/gif"
	"log"
	"os"
	"strconv"

	"github.com/golang/freetype/truetype"
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

// var (
// 	port     = os.Getenv("PORT")
// 	fontfile = os.Getenv("GIFTXTBOT_FONT_FILE")
// )

// Preload all the referables
var font *truetype.Font
var fontSizeCache *giftxt.FontFaceCache

func generateGlobals(fontfile, precache string) {
	font = giftxt.LoadTypeFace(fontfile)

	// If precache string is empty then
	// load it from there, else build it
	if len(precache) != 0 {
		file, ok := os.Open(precache)

		if ok != nil {
			log.Println(ok)
			log.Fatal("could not load precache")
		}

		fontSizeCache, ok = giftxt.LoadFontFaceCache(file)

		if ok != nil {
			log.Println(ok)
			log.Fatal("could not load precache")
		}
	} else {
		fontSizeCache = giftxt.GetFontSizeCrossMap(font, textPortSize, cacheTill)
	}
}

// Generates a new GifText from a given text
func generateGifText(text string) *gif.GIF {
	message := giftxt.NewMessage(text)
	face := fontSizeCache.GetFace(font, len(message.Longest))
	slides := giftxt.RenderMessage(face, message.Words)

	return giftxt.MakeGif(slides, delay)
}

func main() {
	port := flag.Uint("port", 8080, "port to bind giftxt server")
	fontfile := flag.String("fontfile", "", "fontfile to use")
	usePreCache := flag.String("precache", "", "precache to use")
	flag.Parse()

	// Generate all to global values
	generateGlobals(*fontfile, *usePreCache)

	app := iris.Default()

	// Method:   GET;
	// /generate takes a text string and returns a GIF.
	app.Get("/generate", func(ctx context.Context) {
		text := ctx.URLParamDefault("text", "hello world")
		giftxt := generateGifText(text)

		gif.EncodeAll(ctx.ResponseWriter(), giftxt)
		ctx.ContentType("image/gif")
	})

	app.Run(iris.Addr(":" + strconv.Itoa(int(*port))))
}
