package main

import (
	"fmt"
	"github.com/abhinavdahiya/go-messenger-bot"
	"log"
	"net/http"
	"os"
)

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

		respMsg := fmt.Sprintf("You've said: %s", callback.Message.Text)
		msg := mbotapi.NewMessage(respMsg)
		bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
	}
}
