package onec

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ProductOnecCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	result, err := c.subdomainService.Remove(idx)
	if err != nil {
		log.Printf("Fail to delete product with id %d: %v", idx, err)
		return
	}
	outputMessage := ""
	if result {
		outputMessage = "Product is deleted"
	} else {
		outputMessage = "Product isn't deleted"
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMessage,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductOnecCommander.Get: error sending reply message to chat - %v", err)
	}
}
