package bot_server

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
