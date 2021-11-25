package handlers

/*func checkForGuilds(s *discordgo.Session) {
	// Get guilds
	guilds, err := s.UserGuilds(100, "", "")
	if err != nil {
		log.Error(err)
		return
	}

	// For each guild, check if in db and save it if it is not
	for _, guild := range guilds {
		server := models.Server{
			DiscordID: guild.ID,
		}

		if err := repositories.FindServerByDiscordID(&server); err == nil {
			continue
		}
		if err := repositories.PersistServer(&server); err != nil {
			log.Error(err)
		}
	}
}
*/