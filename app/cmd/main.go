package main

import (
	"log"
	"os"
	"os/signal"
	"sellerBot/internal/bot_server"
	"sellerBot/internal/config"
	"sellerBot/internal/logger"
	"syscall"
	"time"
)

var Path string = "/home/coral/yandex_disk/packages/photos"

func main() {

	// env переменные в структуру
	err := config.LoadFromFile(".env")
	if err != nil {
		log.Fatal(err)
	}

	// инициализация логгера
	logger.New(config.Values.LoggerLevel)

	bot_server.RunBot(config.Values.BotToken, config.Values.BotDebug)

	// ожидание сигнала завершения
	awaitQuitSignal()

	// процессы для остановки
	go bot_server.StopBot()

	// ожидание процессов и закрытие main()
	waitingAndReturn()
}

func awaitQuitSignal() {
	logger.Infof("Microservice started. Working until a quit signal is received...")
	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Infof("awaitQuitSignal(): Stopping server...")
}

func waitingAndReturn() {
	time.Sleep(time.Second * 2)
	logger.Info("waitingAndReturn(): The stop timer is completed")
}
