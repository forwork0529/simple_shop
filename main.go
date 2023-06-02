
package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "os"
        "strings"

        tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Path string = "/home/coral/yandex_disk/packages/photos"

func main() {

        // Запуск бота
        file , err := os.Open("./token.txt")
        if err != nil{
                log.Fatal(err)
        }
        token , err := ioutil.ReadAll(file)
        if err != nil{
                log.Fatal(err)
        }
        bot, err := tgbotapi.NewBotAPI(string(token[:len(token)-1]))
        if err != nil {
                log.Fatal(err)
        }

        bot.Debug = false

        log.Printf("Authorized on account %s", bot.Self.UserName)
        log.Printf("Bot startet, dubug mode: %v\n", bot.Debug)

        // Бот запущен

        u := tgbotapi.NewUpdate(0)
        u.Timeout = 60

        updates := bot.GetUpdatesChan(u)


        // updates - канал сообщений получаемый от бота
        for update := range updates {
                if update.Message != nil { // If we got a message
                        if bot.Debug{
                                log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
                        }
                        // Вывод клавиатуры на короткое сообщение или старт
                        if len(update.Message.Text) < 3 || update.Message.Text == "/start" {
                                simpleAnswerTextWithKeyBoard(bot, update, fmt.Sprintf("Здравствуйте %v, надеемся Вам у нас понравится)",update.Message.From.UserName ) )
                                continue
                        }

                        switch update.Message.Text{

                        case "Платья" : informativeAnswer(bot, update, "Вот изделия из категории \"Платья\":", Path + "/dresses" )
                        case "Костюмы" : informativeAnswer(bot, update, "Вот изделия из категории \"Костюмы\":",  Path +"/suits")
                        case "Халаты" : informativeAnswer(bot, update, "Вот изделия из категории \"Халаты\":",  Path +"/bathrobe")
                        case "Кимоно" : informativeAnswer(bot, update, "Вот изделия из категории \"Кимоно\":",  Path +"/kimono")
                        case "Аксессуары" : informativeAnswer(bot, update, "Вот изделия из категории \"Аксессуары\":",  Path +"/accessories")
                        case "Новинки" : informativeAnswer(bot, update, "Вот изделия из категории \"Новинки\":",  Path + "/news")
                        default: simpleAnswerTextWithKeyBoard(bot, update, fmt.Sprintf("В каталоге нет такой категории: %v",update.Message.Text))

                        }
                }
        }
}

var (
        numericKeyboard = tgbotapi.NewReplyKeyboard(
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

func informativeAnswer(bot *tgbotapi.BotAPI, update tgbotapi.Update,  newMsgText, path string)  {

        simpleAnswerText(bot, update, newMsgText)

        fileNames , noFiles := returnFilesInPath(path)

        if noFiles{
                simpleAnswerPhoto(bot, update, Path + "/hangres/hangres.jpg")
                simpleAnswerText(bot, update, "Ассортимент пополняется)..")
        }

        for _, productName := range fileNames{

                simpleAnswerText(bot, update, TrimSuffix(productName))
                simpleAnswerPhoto(bot, update, fmt.Sprintf("%s/%s" ,path, productName))
        }
}

// Функция отправляющая  просто текст сообщения
func simpleAnswerText(bot *tgbotapi.BotAPI, update tgbotapi.Update,  newMsgText string) {
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, newMsgText )
        bot.Send(msg)
}

// Функция отправляющая  просто текст сообщения и выводящая клавиатуру
func simpleAnswerTextWithKeyBoard(bot *tgbotapi.BotAPI, update tgbotapi.Update,  newMsgText string){
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, newMsgText )
        // msg.ReplyToMessageID = update.Message.MessageID
        msg.ReplyMarkup = numericKeyboard
        bot.Send(msg)
}


func simpleAnswerPhoto(bot *tgbotapi.BotAPI, update tgbotapi.Update, photoPath string){
        msgPic := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(photoPath) )
        _, err := bot.Send(msgPic)
        if err != nil{
                log.Printf("cant send image: %v\n", err.Error())
        }
}

// Функция возвращающая список именён файлов в папке и значение говорящее пуст ли список
func returnFilesInPath(path string)([]string, bool){
        files, err := ioutil.ReadDir(path)
        if err != nil{
                log.Printf("Cant read direcory: %v\n", err.Error())
        }
        var fileNames []string
        for _, fileInfo := range files{
                fileNames = append(fileNames, fileInfo.Name())
        }
        if len(fileNames) < 1{
                return fileNames, true
        }
        return fileNames, false

}

func TrimSuffix(toPrintProductName string )string{
        if idx := strings.IndexByte(toPrintProductName, '.'); idx >= 0 {
                toPrintProductName = toPrintProductName[:idx]
        }
        return toPrintProductName
}
