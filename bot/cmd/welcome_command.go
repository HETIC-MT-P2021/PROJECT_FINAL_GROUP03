package commands

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"github.com/JackMaarek/go-bot-utils/enum"
	"github.com/JackMaarek/go-bot-utils/helpers"
	"github.com/JackMaarek/go-bot-utils/models"
)

type WelcomeCommand struct {
	gc *GenericCommand
}

func (command WelcomeCommand) Execute() error {
	params := strings.Split(command.gc.Message.Content, " ")
	log.Info("param: ", params[1])

	// Check if there is an index 3
	if len(params) < 3 {
		log.Error("No command message")
		return nil
	}

	// Check that the message is not empty
	sentence := strings.Join(params[2:], " ")
	log.Info("value sentence : ", sentence)
	if len(sentence) <= 2 {
		log.Error("Message too short")
		return nil
	}
	
	// Si c'est une commande, regarde si la personne est propriétaire ou admin ?
	// if isAdmin, err := helpers.MemberHasPermission(command.gc.Session.(*discordgo.Session), command.gc.Message.GuildID, command.gc.Message.Author.ID, discordgo.PermissionAdministrator); err != nil || !isAdmin {
	// 	// Si pas admin --> niksamer
	// 	log.Error("You are not authorized")
	// 	return nil
	// }
	
	// Si admin et commande existe, on dispatch PAS mais on envoit à l'API (domainAPI) (IsCommandExist)
	payload, err := json.Marshal(models.ChangeWelcomeMessageForm{WelcomeMessage: sentence, DiscordID: command.gc.Message.GuildID})
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = helpers.PerformRequest(enum.ChangeWelcomeMessageRoute, enum.Post, payload)
	if err != nil {
		command.gc.Session.(*discordgo.Session).ChannelMessageSend(command.gc.Message.ChannelID, "Je n'ai pas réussi à changer le message de bienvenue.")
		log.Error(err)
		return err
	}

	_, err = command.gc.Session.ChannelMessageSend(command.gc.Message.ChannelID, fmt.Sprintf("Bravo ! Voici le nouveau message de bienvenue : %s", sentence))
	if err != nil {
		log.Error(err)
	}
	return err
}
