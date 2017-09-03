package main

import (
	"fmt"
	"github.com/abhinavdahiya/go-messenger-bot"
	"net/http"
	"os"
)

func main() {
	accessToken := os.Getenv("ACCESS_TOKEN")
	verifyToken := os.Getenv("VERIFY_TOKEN")
	apiSecret := os.Getenv("API_SECRET")
	bot := mbotapi.NewBotAPI(accessToken, verifyToken, apiSecret)
	callbacks, mux := bot.SetWebhook("/webhook")
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "privkey.pem", mux)

	for callback := range callbacks {
		respMsg := fmt.Sprintf("You've said: %s", callback.Message.Text)
		msg := mbotapi.NewMessage(respMsg)
		bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
	}
}
