package main

import (
 "log"
 "os"

 "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
 // Получаем токен из переменной окружения (безопасно!)
 token := os.Getenv("TELEGRAM_BOT_TOKEN")
 if token == "" {
  log.Fatal("Ошибка: задайте TELEGRAM_BOT_TOKEN в переменных окружения")
 }

 // Создаём бота
 bot, err := tgbotapi.NewBotAPI(token)
 if err != nil {
  log.Panic("Не удалось создать бота:", err)
 }

 bot.Debug = false // в продакшене лучше false
 log.Printf("Бот запущен как @%s", bot.Self.UserName)

 // Конфигурация получения обновлений
 u := tgbotapi.NewUpdate(0)
 u.Timeout = 60 // секунд

 // Канал обновлений
 updates := bot.GetUpdatesChan(u)

 // Основной цикл
 for update := range updates {
  // Игнорируем всё, кроме текстовых сообщений
  if update.Message == nil || update.Message.Text == "" {
   continue
  }

  chatID := update.Message.Chat.ID
  text := update.Message.Text
  user := update.Message.From

  log.Printf("[%s] %s", user.UserName, text)

  // Обработка команд
  if update.Message.IsCommand() {
   handleCommand(bot, chatID, text)
   continue
  }

  // Echo: отправляем то же самое
  msg := tgbotapi.NewMessage(chatID, text)
  msg.ReplyToMessageID = update.Message.MessageID // ответ "в чат"
  bot.Send(msg)
 }
}

func handleCommand(bot *tgbotapi.BotAPI, chatID int64, command string) {
 var response string
 switch command {
 case "/start":
  response = "Привет! Я echo-бот. Напиши что-нибудь — я повторю."
 case "/help":
  response = "Просто напиши текст — я пришлю его обратно.\nПоддерживаю команды: /start, /help"
 default:
  response = "Неизвестная команда. Попробуй /help"
 }

 msg := tgbotapi.NewMessage(chatID, response)
 bot.Send(msg)
}