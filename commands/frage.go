package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the frage command.
	Response: embed :: info
	Permission: all

 */

func CMD_frage(s *discordgo.Session, m *discordgo.MessageCreate)  {

	cmdType := ":speech_balloon: Du hast eine Frage?"

	embed := &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Auch wir sind leider keine Hellseher, um dir also helfen zu können benötigen wir eine konkrete Frage.",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
