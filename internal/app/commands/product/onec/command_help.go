package onec

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ProductOnecCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__product__onec - print list of commands\n"+
			"/list__product__onec - get a list of your entity\n"+
			"/get__product__onec - get a list of your entity\n"+
			"/delete__product__onec - delete an existing entity\n"+
			"/new__product__onec -  create a new entity (not implemented)\n"+
			"/edit__product__onec -  edit a entity (not implemented)\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
