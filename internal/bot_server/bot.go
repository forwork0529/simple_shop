package bot_server

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sellerBot/internal/config"
	"sellerBot/internal/logger"
	"strconv"
)

var stopChan chan struct{} = make(chan struct{}, 0)

type Bot struct {
	*tgbotapi.BotAPI
	updatesChan tgbotapi.UpdatesChannel
}

func newBot(token, debug string) (*Bot, error) {

	// botAPI
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	// botAPI Debug Mod
	botAPI.Debug, err = strconv.ParseBool(debug)
	if err != nil {
		logger.Errorf("strconv.ParseBool(): %v", err)
		botAPI.Debug = false
	}

	// bot update channel config
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := botAPI.GetUpdatesChan(u)

	bot := Bot{
		botAPI,
		updates,
	}

	return &bot, nil
}

func RunBot(token, debug string) error {

	bot, err := newBot(token, debug)
	if err != nil {
		logger.Errorf("RunBot(): %v", err)
	}

	go func() {
		defer logger.Info("reading input messages by bot: STOPPED")
		for {
			select {
			case <-stopChan:
				{
					logger.Info("got stop signal for bot")
					return
				}
			case update := <-bot.updatesChan:
				{
					bot.processUpdate(update)
				}
			}
		}
	}()
	return nil
}

func StopBot() {
	stopChan <- struct{}{}
}

func (b *Bot) processUpdate(update tgbotapi.Update) {
	if update.Message != nil { // If we got a message
		if b.Debug {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		// Вывод клавиатуры на короткое сообщение или старт
		if len(update.Message.Text) < 1 || update.Message.Text == "/start" {
			simpleAnswerTextWithKeyBoard(b.BotAPI, update, fmt.Sprintf("Здравствуйте %v, надеемся Вам у нас понравится)", update.Message.From.UserName))
			return
		}

		switch update.Message.Text {

		case "Платья":
			informativeAnswerLink(b.BotAPI, update, "Вот изделия из категории \"Платья\":", config.Values.YandexPath+`/dresses`)
		case "Костюмы":
			informativeAnswerLink(b.BotAPI, update, "Вот изделия из категории \"Костюмы\":", config.Values.YandexPath+`/suits`)
		case "Свитеры":
			informativeAnswerLink(b.BotAPI, update, "Вот изделия из категории \"свитеры\":", config.Values.YandexPath+`/sweaters`)
		case "Кимоно":
			informativeAnswerLink(b.BotAPI, update, "Вот изделия из категории \"Кимоно\":", config.Values.YandexPath+`/kimono`)
		case "Аксессуары":
			informativeAnswerLink(b.BotAPI, update, "Вот изделия из категории \"Аксессуары\":", config.Values.YandexPath+`/accessories`)
		case "Новинки":
			informativeAnswerLink(b.BotAPI, update, "Вот изделия из категории \"Новинки\":", config.Values.YandexPath+`/news`)
		default:
			simpleAnswerTextWithKeyBoard(b.BotAPI, update, fmt.Sprintf("В каталоге нет такой категории: %v", update.Message.Text))

		}
	}
}
