package onec

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wet_demo_bot/internal/app/path"
)

type CallbackListData struct {
	Cursor int `json:"cursor"`
	Limit  int `json:"limit"`
}

func (c *ProductOnecCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	callbackListData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &callbackListData)
	if err != nil {
		log.Printf("ProductOnecCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	c.SendMessageWithList(callback.Message.Chat.ID, callbackListData)
}
