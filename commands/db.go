package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

/*
	This is the db <scriptname> command.
	Response: embed :: info
	Permission: all

 */

func CMD_db(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := "MySQL Rechte nicht gesetzt"

	scriptName := strings.Fields(m.Content)
	var scriptNameVar string;

	if (len(scriptName) == 1){
		scriptNameVar = "support-pp"
	}else{
		scriptNameVar = scriptName[1]
	}

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Du musst in deiner **config.ini** MySQL Rechte aktivieren\n\n**1.** Öffne deine `config.ini`\n **2.** Füge disen Code am Ender der Datei hinzu:\n ```js\n[Scripts.Privileges] \n\""+scriptNameVar+"\" = [\"db\"]```\n **3.** Starte den Bot neu",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
