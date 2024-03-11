package webappbot

import (
	"context"
	"os"
	"os/signal"
	"time"

	"gitlab.com/deliveos/i18n"
	"gitlab.com/deliveos/telegraph"
	"gitlab.com/deliveos/telegraph/models"
	"gitlab.com/deliveos/telegraph/pkg/logger"
	"gitlab.com/deliveos/webappbot/config"
)

func Run(cfg *config.Config, localizer *i18n.Localizer) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	l := logger.New(cfg.Log.Level)
	l.Info("App was launched...")
	bot, err := telegraph.NewBot(os.Getenv("TELEGRAM_API_TOKEN"), func(options *telegraph.BotOptions) {
		options.IsDebug = true
		options.Buffer = 100
		options.PollTimeout = time.Microsecond
		options.BaseURL = "https://api.telegram.org"
	})
	if err != nil {
		l.Fatal(err)
	}
	var a = models.MenuButtonWebApp{
		Type: "web_app",
		Text: "Dond dell me",
		WebApp: models.WebAppInfo{
			URL: "https://meepott.com:5173",
		},
	}
	var params = make(telegraph.Params)
	params.AddInterface("web_app", a.WebApp)
	params.AddNonEmpty("text", a.Text)
	params.AddNonEmpty("type", a.Type)
	bot.CommandHandler("start", func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
		l.Info(*update.Message.From.LanguageCode)
		params.AddNonZero64("chat_id", update.Message.From.ID)
		bot.MakeRequest("setChatMenuButton", params)

		// bot.SendMessage(update.Message.Chat.ID, localizer.Localize("start_message", *update.Message.From.LanguageCode), telegraph.SendMessageConfing{
		// 	ReplyMarkup: models.NewInlineKeyboardMarkup([][]models.InlineKeyboardButton{
		// 		{
		// 			models.InlineKeyboardButton{
		// 				Text: "qweqwe",
		// 				URL:  ar,
		// 				// WebApp: &models.WebAppInfo{
		// 				// 	URL: *ar,
		// 				// },
		// 			},
		// 		},
		// 	},
		// 	),
		// })
	})

	l.Info("Authorized on account %s", *bot.BotInfo.Username)
	bot.Start(ctx)
}
