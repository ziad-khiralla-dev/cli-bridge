package notifier

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendTelegramNotification(botToken string, chatID int64, message string) error {

	bot, err := tgbotapi.NewBotAPI(botToken)

	if err != nil {
		return fmt.Errorf("failed to initialize bot: %w", err)
	}

	msg := tgbotapi.NewMessage(chatID, message)
	
	_, err = bot.Send(msg)
	
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}