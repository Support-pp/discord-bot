package commands

import (
	"github.com/bwmarrin/discordgo"
	"net/http"
	"io/ioutil"
)

/*
	This is the old command.
	Response: embed :: info
	Permission: all

 */

func CMD_old(s *discordgo.Session, m *discordgo.MessageCreate, language string)  {
	embed := &discordgo.MessageEmbed{}

	cmdType := " Version wird nicht mehr unterstützt"
	versionNr := getVersioFromGitHub()

	embed = &discordgo.MessageEmbed{

		Color:       0x4F545C, // Grey
		Title:     "<:support:386227216159211531> **"+cmdType+"** | Support++",
		Description: ":loudspeaker: Huch, die Version habe ich schon fast vergessen, so alt wie die ist.\nBitte update auf die neuste Version, die neuste Version ist **"+versionNr+"**\n\n :inbox_tray: **Download:**\nhttps://github.com/Support-pp/Support-pp/releases/latest",
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}

func getVersioFromGitHub() string {
	url := "https://raw.githubusercontent.com/Support-pp/Support-pp/main/VERSION"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)

}
