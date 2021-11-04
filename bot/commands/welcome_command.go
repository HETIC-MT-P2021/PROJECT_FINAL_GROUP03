package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/bot/env"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/bot/models"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
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
	
	// // Si c'est une commande, regarde si la personne est propriétaire ou admin ?
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

	client :=  &http.Client{}
	r, err := http.NewRequest("POST", env.GetVariable("DOMAIN_API_URL") + "/commands/change-welcome-message", strings.NewReader(string(payload)))
	if err != nil {
		log.Error(err)
		_, err := command.gc.Session.(*discordgo.Session).ChannelMessageSend(command.gc.Message.ChannelID, "Je n'ai pas réussi à changer le message de bienvenue.")

		if err != nil {
			log.Error("sendMessageErr: ", err)
		}
		return err
	}

	command.gc.Session.ChannelMessageSend(command.gc.Message.ChannelID, fmt.Sprintf("Bravo ! Voici le nouveau message de bienvenue : %s", sentence))

	_, err = client.Do(r)



	return err
}
