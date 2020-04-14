package commands

import "github.com/bwmarrin/discordgo"

/*
	This is the sf command.
	Response: embed :: info
	Permission: all

 */

func CMD_sp(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := "Wie beantrage ich richtig Support?"

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: "Hallo, ich erkläre dir kurz wie du eine Support Anfrage erstellst.\nBitte beachte folgendes:\n\n⤖ Bitte nenne zu erst deine SinusBot Version. Version 1.0\n⤖ Zu dem deine Skript Version. Version 2.5.1\n⤖ Danach dein Betriebssystem auf dem der Bot installiert ist.: Linux Ubuntu 16.04\n⤖ So nun erkläre dein Problem (versuche alle wichtige Details rein zuschreiben.)\n\nBeispiel: Der Parameter vom -\nChannel Edit Module wird bei mir nicht richtig erkannt. Ich versuche es mit /offlineparameter. Was ist falsch?\n\n	Anlagen:\nBitte sende uns folgende Anlagen:\n⤖ Log (Instance log)\n⤖ Bild von deiner Config\n\nSo nun bist du dran. Wenn du alle Punkt berücksichtigst erhältst du schnell eine passende Antwort. :smiley:",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
