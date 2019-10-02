package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the kein support command.
	Response: embed :: info
	Permission: all
 */

func CMD_ksp(s *discordgo.Session, m *discordgo.MessageCreate)  {

	cmdType := "Kein Support"

	embed := &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:       "<:support:386227216159211531> **" + cmdType + "** | Support++",
		Description: "Für diese Software/Script können wir leider keinen Support anbieten.\n**Bitte wende dich an den Entwickler selbst.**\n\n **Software von VerHext**\nLeider kann VerHext nur noch Support für Support++ anbieten, alle anderen Scripts werden nicht mehr unterstützt.\n\n **Sinusbot**\n Wir sind nicht der Offizielle Support von Sinusbot. Bitte stelle alle Fragen rund um den Sinusbot, auf dem Offiziellen Discord: **https://discord.gg/h6s5Ykc**\n\nWir bitten um Verständnis.",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
