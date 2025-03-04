package onec

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ProductOnecCommander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong args", args)
		return
	}

	product, err := c.subdomainService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductOnecCommander.Get: error sending reply message to chat - %v", err)
	}
}
