package main

import (
    "os"
    "strings"
    "fmt"
    "log"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "os/exec"
)

func main() {
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil { // ignore any non-Message Updates
            continue
        }

        switch command := strings.Fields(update.Message.Text)[0]; command {
        case "/uptime":
            cmd := exec.Command("uptime")
            stdout, _ := cmd.Output()
            fmt.Print(string(stdout))
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(stdout))
            msg.ReplyToMessageID = update.Message.MessageID
            bot.Send(msg)
        case "/uname":
            cmd := exec.Command("uname", "-a")
            stdout, _ := cmd.Output()
            fmt.Print(string(stdout))
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(stdout))
            msg.ReplyToMessageID = update.Message.MessageID
            bot.Send(msg)
        case "/cat":
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "meow 🐱")
            msg.ReplyToMessageID = update.Message.MessageID
            bot.Send(msg)
        default:
            fmt.Println(command)
        }
        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
    }
}
