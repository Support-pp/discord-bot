package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the sf command.
	Response: embed :: info
	Permission: all

 */

func CMD_sf(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := "Das Script funktioniert nicht?"

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "** *\"Das Problem sitzt meist vor dem Bildschirm\"* **\nBei dir funktioniert unser Script nicht? Das Problem liegt zu 99% auf deiner Seite.\nÜberprüfe also bitte ob du alle Felder, die mit einem (*) gekennzeichnet sind, ausgefüllt hast.",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
