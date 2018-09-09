package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the config command.
	Response: embed :: info
	Permission: all

 */

func CMD_config(s *discordgo.Session, m *discordgo.MessageCreate)  {

	cmdType := "Einstellungsfehler"

	embed := &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Hallo, du erhältst diese Nachricht, da du nicht in der Lage warst die Beschreibung vom Support++ Script zu lesen!\n\nAber ich sage dir gerne was drinnen steht!\n**1.** Fülle alle Pflichtfelder aus! Diese sind extra mit einem ( * ) markiert.\n**2.** Schreibe in die Felder wie z.B. Servergruppen-IDs keine Wörter oder Sonderzeichen rein",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
