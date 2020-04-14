package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the sinusbot to old command.
	Response: embed :: info
	Permission: all

 */

func CMD_msv(s *discordgo.Session, m *discordgo.MessageCreate, language string) {
	embed := &discordgo.MessageEmbed{}

	cmdType := "SinusBot zu alt"

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:       "<:support:386227216159211531> **" + cmdType + "** | Support++",
		Description: "Seid der Support++ Version 2.5.0 benötigst du mindestens die SinusBot Version 1.0\n\nFür Linux:\nhttps://forum.sinusbot.com/resources/internal-alpha.1/",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
