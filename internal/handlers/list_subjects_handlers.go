package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"gitlab.com/deliveos/i18n"
	"gitlab.com/deliveos/telegraph"
	"gitlab.com/deliveos/telegraph/models"
	"gitlab.com/deliveos/telegraph/pkg/logger"
	"gitlab.com/deliveos/webappbot/internal/sqlc"
	"gitlab.com/deliveos/webappbot/pkg/linq"
)

func ListSubjects(logger *logger.Logger, localizer *i18n.Localizer) telegraph.HandlerFunc {
	return func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
		conn, err := sql.Open("postgres", os.Getenv("PG_URL")+"?sslmode=disable")
		if err != nil {
			logger.Error(err)
		}
		q := sqlc.New(conn)
		subjects, err := q.GetAllSubjects(ctx)
		if err != nil {
			logger.Error(err)
		}
		q.Close()
		var mrk [][]models.InlineKeyboardButton
		for i := 0; i < len(subjects); i++ {
			mrk = append(mrk, []models.InlineKeyboardButton{
				*bot.NewInlineKeyboardButtonWithCallback(subjects[i].Name.String+" | "+subjects[i].Teacher.String+" | "+subjects[i].Classrom.String, string(subjects[i].ID), func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
					logger.Info("%d", update.CallbackQuery.From.ID)
					bot.SendMessage(update.Message.From.ID, *update.CallbackQuery.Data, telegraph.SendMessageConfing{
						ReplyMarkup: *models.NewInlineKeyboardMarkup([][]models.InlineKeyboardButton{
							{
								*bot.NewInlineKeyboardButtonWithCallback(localizer.Localize("edit", *update.CallbackQuery.From.LanguageCode), *update.CallbackQuery.Data, func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {}),
								*bot.NewInlineKeyboardButtonWithCallback(localizer.Localize("delete", *update.CallbackQuery.From.LanguageCode), *update.CallbackQuery.Data, func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {}),
							},
						}),
					})
				}),
			})
		}

		bot.SendMessage(update.Message.From.ID, fmt.Sprintf("%s:\n %s", localizer.Localize("list_of_your_subjects", *update.Message.From.LanguageCode), strings.Join(linq.Map(subjects, func(subject sqlc.SchedulesSubject) string { return subject.Name.String }), "\n")), telegraph.SendMessageConfing{
			ReplyMarkup: *models.NewInlineKeyboardMarkup(mrk),
		})
	}
}
