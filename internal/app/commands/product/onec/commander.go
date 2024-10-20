package onec

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wet_demo_bot/internal/app/path"
	service "github.com/wet_demo_bot/internal/service/product/onec"
)

type OnecCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented

	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath)
}

type ProductOnecCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService service.OnecService
}

func NewProductOnecCommander(bot *tgbotapi.BotAPI) *ProductOnecCommander {

	subdomainService := service.NewDummyOnecService()

	return &ProductOnecCommander{
		bot:              bot,
		subdomainService: subdomainService,
	}

}

func (c *ProductOnecCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ProductOnecCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
