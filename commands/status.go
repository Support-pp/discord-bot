package commands

import (
	"github.com/bwmarrin/discordgo"
	"os"
)

/*
	This is the status command.
	Response: embed :: info
	Permission: all

 */

func CMD_status(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := "Status"


	tag, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Docker container tag: `"+tag+"`",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
