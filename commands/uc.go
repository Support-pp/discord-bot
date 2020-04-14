package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the uc command.
	Response: embed :: info
	Permission: all

 */

func CMD_uc(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := "Unvollständige Config"

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:       "<:support:386227216159211531> **" + cmdType + "** | Support++",
		Description: "**Leider ist das nicht die vollständige Script Config!**\nUm dir helfen zu können benötigen wir Screenshots von der gesamten Script Config.",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
