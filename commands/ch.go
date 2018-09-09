package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the channel edit command.
	Response: embed :: info
	Permission: all

 */

func CMD_ch(s *discordgo.Session, m *discordgo.MessageCreate)  {

	cmdType := "Channel Edit"

	embed := &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "**ChannelEdit**: Der Parameter ist ein Pflichtfeld und muss immer angegeben werden. Durch diesen Parameter lassen sich die Channel alle einzeln öffnen / schließen.\n\nZum Beispiel nehmen wir den Parameter \"**ts**\". Nun kann ich den Support Channel öffnen, in dem ich den Bot mit der Nachricht \"**!online ts**\" anschreibe.\n\n Dieser Parameter wird auch für das Time Module benötigt. Möchtest du alle Support Channel Öffnen / Schließen. Trage zusätzlich deine Supporter Id's in das Feld \"Default Supporter\" ein. Diese können alle Channel mit dem Command \"**!online**\" bzw. \"**!offline**\" editieren. Der Parameter muss für jeden Supportraum anders sein. Und sollte aus einem Wort bestehen. Ich hoffe hiermit ist deine Frage beantwortet :smiley:",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
