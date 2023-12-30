package bot_server

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"sellerBot/internal/logger"
	"sellerBot/internal/models"
)

func informativeAnswerLink(bot *tgbotapi.BotAPI, update tgbotapi.Update, newMsgText, path string) {

	simpleAnswerText(bot, update, newMsgText)

	fileNames, noFiles := returnFilesInPath(path)

	if noFiles {
		simpleAnswerPhoto(bot, update, `./packages/photos/hangres/hangres.jpg`)
		simpleAnswerText(bot, update, "Ассортимент пополняется)..")
	}

	for _, productName := range fileNames {

		simpleAnswerText(bot, update, TrimSuffix(productName))
		readedBytes, err := os.ReadFile(path + `\` + productName)
		if err != nil {
			logger.Errorf("os.ReadFile(%v%v)", path, productName)
			continue
		}
		simpleAnswerText(bot, update, string(readedBytes))
	}
}

func informativeAnswerPhoto(bot *tgbotapi.BotAPI, update tgbotapi.Update, newMsgText, path string) {

	simpleAnswerText(bot, update, newMsgText)

	fileNames, noFiles := returnFilesInPath(path)

	if noFiles {
		simpleAnswerPhoto(bot, update, `./packages/photos/hangres/hangres.jpg`)
		simpleAnswerText(bot, update, "Ассортимент пополняется)..")
	}

	for _, productName := range fileNames {

		simpleAnswerText(bot, update, TrimSuffix(productName))
		simpleAnswerPhoto(bot, update, fmt.Sprintf("%s/%s", path, productName))
	}
}

// Функция отправляющая  просто текст сообщения
func simpleAnswerText(bot *tgbotapi.BotAPI, update tgbotapi.Update, newMsgText string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, newMsgText)
	bot.Send(msg)
}

// Функция отправляющая  просто текст сообщения и выводящая клавиатуру
func simpleAnswerTextWithKeyBoard(bot *tgbotapi.BotAPI, update tgbotapi.Update, newMsgText string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, newMsgText)
	// msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = models.NumericKeyboard
	bot.Send(msg)
}

func simpleAnswerPhoto(bot *tgbotapi.BotAPI, update tgbotapi.Update, photoPath string) {
	msgPic := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(photoPath))
	_, err := bot.Send(msgPic)
	if err != nil {
		log.Printf("cant send image: %v\n", err.Error())
	}
}
