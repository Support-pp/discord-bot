package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the pro command.
	Response: embed :: info
	Permission: all

 */

func CMD_pro(s *discordgo.Session, m *discordgo.MessageCreate)  {

	cmdType := "Spende um ein PRO Mitglied zu werden"

	embed := &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Die PRO Version beinhaltet folgende Funktionen:\n\n:small_blue_diamond: Beantworte Tickets direkt über Discord (**/reply**)\n:small_blue_diamond: Bekomme eine Datenbank zur verfügung gestellt\n:small_blue_diamond: Erwähung im SourceCode und im Dashboard\n:small_blue_diamond:  Unterstütze die Weiterentwicklung von Support++ :heart:\n\nSpende uns nun unter **https://support-pp.de/pro**",
		Image: &discordgo.MessageEmbedImage{
			URL: "https://media.discordapp.net/attachments/348417474905243651/488377328322347008/pro.png",
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
