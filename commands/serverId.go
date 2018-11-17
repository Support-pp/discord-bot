package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the serverid command.
	Response: embed :: info
	Permission: all
 */

func CMD_serverId(s *discordgo.Session, m *discordgo.MessageCreate)  {

	cmdType := "ServerId"

	embed := &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Du kannst die ServerId ganz einfach heruasfinden. Discord hat eine Schritt-für-Schritt Anleitung erstellt. https://support.discordapp.com/hc/de/articles/206346498-Wie-finde-ich-meine-Server-ID-",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
