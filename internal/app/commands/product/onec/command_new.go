package onec

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wet_demo_bot/internal/model/product/onec"
)

func (c *ProductOnecCommander) New(inputMessage *tgbotapi.Message) {

	outputMessage := ""
	_, err := c.subdomainService.Create(onec.Onec{})

	if err != nil {

		if err != nil {
			outputMessage = "Command New isn't implemented"
			log.Println(outputMessage)
		}

	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMessage,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductOnecCommander.Edit: error sending reply message to chat - %v", err)
	}
}
