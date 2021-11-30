package cron

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/services"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"time"
)

// InitBirthdayReminderJob set up a cron job to check if it's someone's birthday today
func InitBirthdayReminderJob(session *discordgo.Session) {
	task := gocron.NewScheduler(time.UTC)
	if _, err := task.Every(1).Day().At("13:00").Do(services.RemindBirthdays, session); err != nil {
		log.Warn("could not launch birthday reminder job")
	}
	task.StartAsync()
}
