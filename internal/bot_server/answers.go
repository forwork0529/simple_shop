package bot_server

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"sellerBot/internal/logger"
	"sellerBot/internal/models"
)

func informativeAnswerLink(bot *tgbotapi.BotAPI, update tgbotapi.Update, newMsgText, path, folder string) {

	simpleAnswerText(bot, update, newMsgText)

	fileNames, noFiles := returnFilesInPath(path + `/` + folder)

	if noFiles {
		simpleAnswerPhoto(bot, update, `./packages/photos/hangres/hangres.jpg`)
		simpleAnswerText(bot, update, "Ассортимент пополняется)..")
	}

	shopIdx := retIndexSymbol(folder)
	var countIdx = 1

	for _, productName := range fileNames {

		simpleAnswerText(bot, update, TrimSuffix(fmt.Sprintf("Изделие #%s%v", shopIdx, countIdx)))
		readedBytes, err := os.ReadFile(path + folder + `/` + productName)
		if err != nil {
			logger.Errorf("os.ReadFile(%v/%v): %v", path, productName, err)
			continue
		}
		simpleAnswerText(bot, update, string(readedBytes))
		countIdx += 1
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
		logger.Errorf("cant send image: %v\n", err.Error())
	}
}
