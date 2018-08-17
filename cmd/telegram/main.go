package main

import (
	"flag"
	"log"
	"net/url"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	tgBotToken := flag.String("token", "", "telegram bot access token")
	tgGifTxtBotServer := flag.String("server", "", "giftxt base endpoint url")
	flag.Parse()

	bot, ok := tb.NewBot(tb.Settings{
		Token:  *tgBotToken,
		Poller: &tb.LongPoller{Timeout: 2 * time.Second},
	})

	if ok != nil {
		log.Fatal("Bot initialization failed", ok)
	}

	log.Println("Bot Setup Successfully")

	bot.Handle(tb.OnQuery, func(q *tb.Query) {
		log.Println(q.ID, "Received query request")

		// Get the URL and add parameter
		tgGifTxtBotServerURL, _ := url.Parse(*tgGifTxtBotServer)
		tgGifTxtBotServerURL.Query().Add("text", q.Text)

		// Create inline request response
		result := &tb.GifResult{
			URL:      tgGifTxtBotServerURL.String(),
			ThumbURL: tgGifTxtBotServerURL.String(),
		}

		result.SetResultID("result-medium-speed")

		// Dispatch the result response
		ok = bot.Answer(q, &tb.QueryResponse{
			Results: []tb.Result{
				result,
			},
		})

		if ok != nil {
			log.Println(q.ID, "Failed serving to query")
		}

		log.Println(q.ID, "Serving request complete")
	})

	bot.Start()
}
