package bot

import (
	"context"
	"strings"

	"github.com/d1nch8g/swap/gen/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	Host     string
	Database *database.Queries
	Bot      *bot.Bot
}

func RunBot(host, token string, db *database.Queries) error {
	mybot := &Bot{
		Host:     host,
		Database: db,
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(mybot.defaultHandler),
	}
	b, err := bot.New(token, opts...)
	if err != nil {
		return err
	}
	mybot.Bot = b

	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, mybot.helpHandler)

	b.Start(context.Background())
	return nil
}

func (s *Bot) helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: `Для получения помощи по заявке отправьте пожалуйста в ответном сообщении по форме

1) Вашу электронная почта
2) Номер заявки (если известен)
3) Подробное описание проблемы с деталями для администрации сайта

Сообщение будет доставлено до администрации сайта, ответ прийдет через бота и на email.`,
	})
}

func (s *Bot) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if strings.HasPrefix(update.Message.Text, "1)") {
		_, err := s.Database.CreateBotMessage(ctx, database.CreateBotMessageParams{
			UserID:  nil,
			OrderID: nil,
			Message: update.Message.Text,
			Checked: false,
		})
		if err != nil {
			logrus.Error("unbale to deliver bot message: ", err)
			return
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   `Ваше сообщение доставлено, ожидайте ответа в этом чате`,
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: `Введена непонятная команда, список допустимых комманд:

/help`,
	})

}
