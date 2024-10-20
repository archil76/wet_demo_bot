package product

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wet_demo_bot/internal/app/commands/product/onec"
	"github.com/wet_demo_bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ProductCommander struct {
	bot           *tgbotapi.BotAPI
	onecCommander onec.OnecCommander
}

func NewProductCommander(bot *tgbotapi.BotAPI) *ProductCommander {
	return &ProductCommander{bot: bot, onecCommander: onec.NewProductOnecCommander(bot)}
}

func (c *ProductCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "onec":
		c.onecCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ProductCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ProductCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "onec":
		c.onecCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ProductCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
