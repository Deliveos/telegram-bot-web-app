package keyboards

import (
	"gitlab.com/deliveos/i18n"
	"gitlab.com/deliveos/telegraph"
	"gitlab.com/deliveos/telegraph/models"
	"gitlab.com/deliveos/webappbot/internal/handlers"
)

type Weekday string

const (
	Monday   Weekday = "Monday"
	Tuestay  Weekday = "Tuestay"
	Wendsday Weekday = "Wendsday"
	Thursday Weekday = "Thursday"
	Friday   Weekday = "Friday"
	Saturday Weekday = "Saturday"
	Sunday   Weekday = "Sunday"
)

func GetWeekdaysKeyboard(localizer i18n.Localizer, bot telegraph.Bot, message *models.Message) models.InlineKeyboardMarkup {
	return *models.NewInlineKeyboardMarkup([][]models.InlineKeyboardButton{
		{
			*bot.NewInlineKeyboardButtonWithCallback(localizer.Localize("Monday", *message.From.LanguageCode), string(Monday), handlers.AckTime(localizer)),
			*bot.NewInlineKeyboardButtonWithCallback(localizer.Localize("Tuestay", *message.From.LanguageCode), string(Tuestay), handlers.AckTime(localizer)),
			*bot.NewInlineKeyboardButtonWithCallback(localizer.Localize("Wendsday", *message.From.LanguageCode), string(Wendsday), handlers.AckTime(localizer)),
		},
		{
			*bot.NewInlineKeyboardButtonWithCallback(localizer.Localize("Thursday", *message.From.LanguageCode), string(Thursday), handlers.AckTime(localizer)),
			*bot.NewInlineKeyboardButtonWithCallback(localizer.Localize("Friday", *message.From.LanguageCode), string(Friday), handlers.AckTime(localizer)),
			*bot.NewInlineKeyboardButtonWithCallback(localizer.Localize("Saturday", *message.From.LanguageCode), string(Saturday), handlers.AckTime(localizer)),
		},
		{
			*bot.NewInlineKeyboardButtonWithCallback(localizer.Localize("Sunday", *message.From.LanguageCode), string(Sunday), handlers.AckTime(localizer)),
		},
	})
}
