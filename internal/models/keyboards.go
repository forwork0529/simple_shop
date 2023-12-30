package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	NumericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Платья"),
			tgbotapi.NewKeyboardButton("Костюмы"),
			tgbotapi.NewKeyboardButton("Халаты"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Кимоно"),
			tgbotapi.NewKeyboardButton("Аксессуары"),
			tgbotapi.NewKeyboardButton("Новинки"),
		),
	)
)
