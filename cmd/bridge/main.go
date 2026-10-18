package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/ziad-khiralla-dev/cli-bridge/internal/executor"
	"github.com/ziad-khiralla-dev/cli-bridge/internal/notifier"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a command to execute")
		fmt.Println("Usage example: go run cmd/bridge/main.go \"dir\"")
		return
	}

	command := os.Args[1]

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, falling back to system environment variables")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" {
		log.Fatal("Error: TELEGRAM_BOT_TOKEN not found")
	}
	
	if chatIDStr == "" {
		log.Fatal("Error: TELEGRAM_CHAT_ID not found")
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Error: Invalid TELEGRAM_CHAT_ID: %v", err)
	}

	fmt.Printf("Executing command: %s ...\n", command)

	output, duration, err := executor.ExecuteCommand(command)
	
	var statusMsg string
	if err != nil {
		statusMsg = fmt.Sprintf("[ Command Execution Failed! ]\n\n --> Command: ' %s ' \n --> Error: ' %v ' \n --> Duration: ' %v '", command, err, duration)
	} else {
		statusMsg = fmt.Sprintf("[ Command Executed Successfully! ]\n\n --> Command: ' %s ' \n --> Output: ' %s ' \n --> Duration: ' %v '", command, output, duration)
	}
	fmt.Println("Sending notification to Telegram ...")
	err = notifier.SendTelegramNotification(botToken, chatID, statusMsg)
	if err != nil {
		fmt.Printf("Failed to send notification: %v\n", err)
	} else {
		fmt.Println("Notification sent successfully! Task Completed!")
	}
}