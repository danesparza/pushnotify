package main

import (
	"flag"
	"log"

	"github.com/gregdel/pushover"
)

//	Set up our flags
var (
	apiKey   = flag.String("apiKey", "adu3foje2vrme6dbqu5zs2dawi5e2o", "The pushover API key to send from")
	to       = flag.String("to", "", "Pushover recipient")
	priority = flag.Int("priority", 0, "The message priority")
	title    = flag.String("title", "", "The message title (optional)")
	message  = flag.String("message", "", "The message to send")
	url      = flag.String("url", "", "The url to send (optional)")
	urlTitle = flag.String("urlTitle", "", "The title to use for the url (optional)")
	sound    = flag.String("sound", "pushover", "The sound to use. One of: pushover, bike, bugle, cashregister, classical, cosmic, falling, gamelan, incoming, intermission, magic, mechanical, pianobar, siren, spacealarm, tugboat, alien, climb, persistent, echo, updown, none")
)

func main() {
	//	Parse the flags
	flag.Parse()

	//	Create our client, recipient and message
	pushClient := pushover.New(*apiKey)
	recipient := pushover.NewRecipient(*to)
	message := &pushover.Message{
		Message:  *message,
		Title:    *title,
		Priority: *priority,
		URL:      *url,
		URLTitle: *urlTitle,
		Sound:    *sound,
	}

	// Send the message to the recipient
	_, err := pushClient.SendMessage(message, recipient)
	if err != nil {
		log.Printf("Error sending pushover message: %v\n", err)
	}
}
