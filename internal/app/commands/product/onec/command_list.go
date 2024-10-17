package onec

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wet_demo_bot/internal/app/path"
)

func (c *ProductOnecCommander) List(inputMessage *tgbotapi.Message) {
	var cursor, limit int
	if limit == 0 {
		limit = 4
	}

	c.SendMessageWithList(inputMessage.Chat.ID, CallbackListData{Cursor: cursor, Limit: limit})

}

func (c *ProductOnecCommander) SendMessageWithList(chatID int64, callbackListData CallbackListData) {

	outputMsgText := "Here all the products: \n\n"

	products, _ := c.subdomainService.List(callbackListData.Cursor, callbackListData.Limit)
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	if len(products) == 0 {
		outputMsgText = "It's all"
		c.SendSimleMessage(chatID, outputMsgText)
		return

	}
	msg := tgbotapi.NewMessage(chatID, outputMsgText)

	callbackListData.Cursor = callbackListData.Cursor + callbackListData.Limit

	serializedData, _ := json.Marshal(callbackListData)

	callbackPath := path.CallbackPath{
		Domain:       "product",
		Subdomain:    "onec",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductOnecCommander.List: error sending reply message to chat - %v", err)
	}

}

func (c *ProductOnecCommander) SendSimleMessage(chatID int64, msgText string) {

	msg := tgbotapi.NewMessage(chatID, msgText)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductOnecCommander.Simle: error sending simple message to chat - %v", err)
	}

}
