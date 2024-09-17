package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/d1nch8g/swap/gen/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/labstack/echo/v4"
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

	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, mybot.helloHandler)

	b.Start(context.Background())
	return nil
}

func (s *Bot) helloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		ParseMode: models.ParseModeMarkdown,
		Text: "Здравствуйте, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*" + `

Отправьте пожалуйста:
1) Номер заявки
2) Ваш email
3) Сообщение которое хотите доставить до службы поддержки

`,
	})
}

func (s *Bot) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	switch {
	case !strings.HasPrefix(update.Message.Text, "/"):

		_, err := s.Database.CreateBotMessage(ctx, database.CreateBotMessageParams{
			UserID:  new(int64),
			OrderID: new(int64),
			Message: update.Message.Text,
			Checked: false,
		})
		if err != nil {
			echo.New().Logger.Errorf("unable to create bot message")
		}
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Ваше сообщение доставлено, ожидайте ответа в этом чате",
		})

	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text: fmt.Sprintf(`Здравствуйте, это бот обменника %s.
	
	Получить помощь по сделке: /help`, s.Host),
		})
	}

}
