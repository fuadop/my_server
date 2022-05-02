package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fuadop/my_server/internal/handlers"
	"github.com/fuadop/sendchamp"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	publicKey := os.Getenv("SENDCHAMP_PUBLIC_KEY")
	client := sendchamp.NewClient(publicKey, sendchamp.ModeLive)

	mux := http.NewServeMux()

	mux.Handle("/", handlers.GetIndex)
	mux.Handle("/sms", handlers.SendMessage(client))
	mux.Handle("/voice", handlers.SendVoice(client))

	log.Println("Listening on port 5000")
	http.ListenAndServe(":5000", mux)
}
