package bot

import (
	"cryptoApi/pkg/auth"
	"cryptoApi/pkg/ethereum"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Start() {
	apiKey := os.Getenv("TELEGRAM_BOT_API_KEY")

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatalf("Error initializing bot API: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates channel: %v", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		user := update.Message.From
		exists := auth.CheckUserExists(user.ID)
		if !exists {
			dto := auth.CreateUserDTO{
				ID:           user.ID,
				IsBot:        user.IsBot,
				FirstName:    user.FirstName,
				LastName:     user.LastName,
				UserName:     user.UserName,
				LanguageCode: user.LanguageCode,
			}
			auth.AddUserToDatabase(dto)
		}

		log.Printf("[%s] %s", user.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = "Привет! Я бот, который поможет тебе узнать цену Ethereum."
		case "help":
			msg.Text = "Команды:\n/start - начать общение с ботом\n/eth - узнать цену Ethereum"
		case "eth":
			ethereum.GetPrice()
		default:
			msg.Text = "Я не знаю эту команду :("
		}

		bot.Send(msg)
	}
}
