package main

import (
	"fmt"
	"github.com/FlashBoys/go-finance"
	"github.com/abhinavdahiya/go-messenger-bot"
	"github.com/davecgh/go-spew/spew"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type QuoteResponse struct {
}

func main() {
	accessToken := os.Getenv("ACCESS_TOKEN")
	verifyToken := os.Getenv("VERIFY_TOKEN")
	apiSecret := os.Getenv("API_SECRET")
	certPath := os.Getenv("CERT_PATH")
	keyPath := os.Getenv("KEY_PATH")
	bot := mbotapi.NewBotAPI(accessToken, verifyToken, apiSecret)
	callbacks, mux := bot.SetWebhook("/webhook")
	go http.ListenAndServeTLS(":8443", certPath, keyPath, mux)

	for callback := range callbacks {
		log.Printf("[%#v] %s", callback.Sender, callback.Message.Text)

		greet := mbotapi.NewMessage("Hello")
		bot.Send(callback.Sender, greet, mbotapi.RegularNotif)

		var respMsg string
		if _, err := strconv.ParseInt(callback.Message.Text, 10, 64); err != nil {
			respMsg = "Hello! This is a lab experiment. A Hong Kong stock quoting bot. Please provide stock number *NUMBER ONLY* to quote your stock."
		} else {
			quoteResp, err := finance.GetQuote(fmt.Sprintf("%s.HK", strings.TrimSpace(callback.Message.Text)))
			if err != nil {
				log.Printf("Failed to quote stock [%s]. Error: %v\n", callback.Message.Text, err)
			}
			respMsg = spew.Sdump(quoteResp)
		}
		msg := mbotapi.NewMessage(respMsg)
		bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
	}
}
