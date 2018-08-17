package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/ksdme/giftxt"
)

func main() {
	fontfile := flag.String("fontfile", "", "fontfile to precalculate font sizes")
	till := flag.Int("till", 30, "cache size in terms of character length limit")
	size := flag.Int("size", 275, "size to use while calculating")

	flag.Parse()

	if *fontfile == "" {
		log.Fatal("no fontfile found")
	}

	absFontFile, ok := filepath.Abs(*fontfile)

	if ok != nil {
		log.Fatal("unable to build abs path")
	}

	font := giftxt.LoadTypeFace(absFontFile)
	cache := giftxt.GetFontSizeCrossMap(font, int32(*size), *till)

	gobfile := absFontFile + ".json"
	gob, ok := os.Create(gobfile)
	defer gob.Close()

	if ok != nil {
		log.Println(ok)
		log.Fatal("cache file could not be created")
	}

	ok = cache.Export(gob)

	if ok != nil {
		log.Println(ok)
		log.Fatal("cache could not be serialized")
	}
}
