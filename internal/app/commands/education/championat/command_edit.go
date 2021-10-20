package championat

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *ChampionatCommander) Edit(inputMessage *tgbotapi.Message) {
	msgParts := strings.SplitN(inputMessage.Text, " ", 2)
	editData := ChampionatEditData{}

	err := json.Unmarshal([]byte(msgParts[1]), &editData)
	if err != nil {
		log.Printf("fail to unmarshal championat from text: %v. Error: %v", msgParts[1], err)
		return
	}

	err = c.championatService.Edit(editData.ID, editData.Title)
	if err != nil {
		log.Printf("fail to edit championat with id %v: %v", editData.ID, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		fmt.Sprintf("Championat with id %v was edited!", editData.ID),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ChampionatCommander.Help: error sending reply message to chat - %v", err)
	}
}
