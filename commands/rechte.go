package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the rechte command.
	Response: embed :: info
	Permission: all

 */

func CMD_rechte(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := "Bot Rechte"

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Beachte bitte das der Bot folgende Rechte benötigt:\nb_channel_create_with_password\ni_channel_modify_power\nb_channel_modify_description\nb_channel_modify_name\nb_channel_modify_password\n\n**Dieses Recht benötigt der Bot nicht!** *(Sonst editiert er nur offline)*\nb_channel_create_modify_with_force_password",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
