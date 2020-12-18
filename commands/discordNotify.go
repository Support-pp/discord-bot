package commands

import (
	"github.com/bwmarrin/discordgo"
)

/*
	This is the discordnotifydocs command.
	Response: embed :: info
	Permission: all
 */

func CMD_discordNotify(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := " Anleitung zum Einrichten der Discord Benachrichtigung"

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Hallo, du findest eine Anleitung zum Einrichten [hier](https://docs.support-pp.de/de/latest/modules/discord.html)",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}

