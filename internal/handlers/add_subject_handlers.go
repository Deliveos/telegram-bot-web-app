package handlers

import (
	"context"
	"database/sql"
	"os"
	"regexp"
	"time"

	"gitlab.com/deliveos/i18n"
	"gitlab.com/deliveos/telegraph"
	"gitlab.com/deliveos/telegraph/pkg/logger"
	"gitlab.com/deliveos/webappbot/internal/sqlc"
)

// Starts subject addition dialog
func AskSubjectName(logger *logger.Logger, localizer *i18n.Localizer, subject *sqlc.SchedulesSubject) telegraph.HandlerFunc {
	return func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
		logger.Info(*update.Message.Text)
		bot.SendMessageWithCallback(update.Message.Chat.ID, localizer.Localize("send_me_subject_name", *update.Message.From.LanguageCode), askSubjectNameHandler(logger, localizer, subject))
	}
}

func askSubjectNameHandler(logger *logger.Logger, localizer *i18n.Localizer, subject *sqlc.SchedulesSubject) telegraph.HandlerFunc {
	return func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
		logger.Info(*update.Message.Text)
		if *update.Message.Text == "" {
			bot.SendMessageWithCallback(update.Message.Chat.ID, localizer.Localize("invalid_subject_name", *update.Message.From.LanguageCode), askSubjectNameHandler(logger, localizer, subject))
			return
		}
		subject.Name = sql.NullString{
			String: *update.Message.Text,
			Valid:  true,
		}
		bot.SendMessageWithCallback(update.Message.Chat.ID, localizer.Localize("send_me_teacher_name", *update.Message.From.LanguageCode), askTeatherNameHandler(logger, localizer, subject))
	}
}

func askTeatherNameHandler(logger *logger.Logger, localizer *i18n.Localizer, subject *sqlc.SchedulesSubject) telegraph.HandlerFunc {
	return func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
		logger.Info(*update.Message.Text)
		if *update.Message.Text == "" {
			bot.SendMessageWithCallback(update.Message.Chat.ID, localizer.Localize("invalid_teacher_name", *update.Message.From.LanguageCode), askTeatherNameHandler(logger, localizer, subject))
			return
		}
		subject.Teacher = sql.NullString{
			String: *update.Message.Text,
			Valid:  true,
		}
		bot.SendMessageWithCallback(update.Message.Chat.ID, localizer.Localize("send_me_classroom_number", *update.Message.From.LanguageCode), askClassroomHandler(logger, localizer, subject))
	}
}

func askClassroomHandler(logger *logger.Logger, localizer *i18n.Localizer, subject *sqlc.SchedulesSubject) telegraph.HandlerFunc {
	return func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
		logger.Info(*update.Message.Text)
		if *update.Message.Text == "" {
			bot.SendMessageWithCallback(update.Message.Chat.ID, localizer.Localize("invalid_classroom_number", *update.Message.From.LanguageCode), askClassroomHandler(logger, localizer, subject))
			return
		}
		subject.Classrom = sql.NullString{
			String: *update.Message.Text,
			Valid:  true,
		}
		conn, err := sql.Open("postgres", os.Getenv("PG_URL")+"?sslmode=disable")
		if err != nil {
			logger.Error(err)
		}
		q := sqlc.New(conn)
		_, err = q.CreateSubject(ctx, sqlc.CreateSubjectParams{
			Name:     subject.Name,
			Teacher:  subject.Teacher,
			Classrom: subject.Classrom,
		})
		q.Close()
		if err != nil {
			logger.Error(err)
			bot.SendMessage(update.Message.Chat.ID, localizer.Localize("failed_to_add_subject", *update.Message.From.LanguageCode), telegraph.SendMessageConfing{})
		} else {
			bot.SendMessage(update.Message.Chat.ID, localizer.Localize("subject_added", *update.Message.From.LanguageCode), telegraph.SendMessageConfing{})
		}
	}
}

func AckTime(l i18n.Localizer) telegraph.HandlerFunc {
	return func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
		bot.SendMessageWithCallback(update.CallbackQuery.From.ID, l.Localize("send_me_time", *update.CallbackQuery.From.LanguageCode), handleTime(l))
	}
}

func handleTime(l i18n.Localizer) telegraph.HandlerFunc {
	return func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
		matched, _ := regexp.MatchString("[0-9]{2}:[0-9]{2}", *update.Message.Text)
		if !matched {
			_, err := time.Parse("HH:mm", *update.Message.Text)
			if err != nil {
				bot.SendMessageWithCallback(update.Message.Chat.ID, "The time is not correct. Please send me time in format HH:mm", handleTime(l))
				return
			}
		}
		bot.SendMessageWithCallback(update.Message.Chat.ID, "Please send me subject name", func(ctx context.Context, bot *telegraph.Bot, update *telegraph.Update) {
			bot.SendMessage(update.Message.Chat.ID, "Subject was successufully adeed =)", telegraph.SendMessageConfing{})
		})
	}
}
