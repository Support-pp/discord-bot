package commands

import (
	"github.com/bwmarrin/discordgo"
)

/*
	This is the passwordsecurity info command.
	Response: embed :: info
	Permission: all
 */

func CMD_tokenAbuse(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := " Sicherheitshinweis | Token abuse"

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: ":loudspeaker: Wir haben deinen gesendeten Token gelöscht. Durch die Veröffentlichung, können andere Zugriff auf deinen Channel erlangen. Gebe diesen Token niemals an andere personen weiter!. Unser Team wird niemals nach deinem Token fragen.",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
