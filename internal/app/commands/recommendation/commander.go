package recommendation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/recommendation/production"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type RecommendationCommander struct {
	bot                 *tgbotapi.BotAPI
	productionCommander Commander
}

func NewRecommendationCommander(bot *tgbotapi.BotAPI) *RecommendationCommander {
	return &RecommendationCommander{
		bot: bot,
		// productionCommander
		productionCommander: production.NewRecommendationProductionCommander(bot),
	}
}

func (c *RecommendationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "production":
		c.productionCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("RecommendationCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *RecommendationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "production":
		c.productionCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("RecommendationCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
