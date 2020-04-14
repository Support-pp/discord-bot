package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the uv command.
	Response: embed :: info
	Permission: all

 */

func CMD_uv(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := "Distribution unbekannt!"

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Uns ist keine Linux Distribution bzw. Version mit diesem Namen bekannt.\n*Auf dem Bild findest du eine Auswahl von Distributionen.*",
		Image: &discordgo.MessageEmbedImage{
			URL: "https://cdn.discordapp.com/attachments/316588859448360962/432544428612124674/Linux_Distribution_Timeline.png",
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
