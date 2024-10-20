package onec

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wet_demo_bot/internal/model/product/onec"
)

func (c *ProductOnecCommander) Edit(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong args", args)
		return
	}

	outputMessage := ""
	err = c.subdomainService.Update(uint64(idx), onec.Onec{})

	if err != nil {
		outputMessage = "Command Edit isn't implemented"
		log.Println(outputMessage)
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
