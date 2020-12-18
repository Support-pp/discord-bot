package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/Support-pp/discord-bot/commands"
	"strings"
)


const EN_CHANNEL = "309033097180217349";
const DE_CHANNEL = "303664307667730440";
var language = "de";

func main() {

	if (os.Getenv("TOKEN") == ""){
		fmt.Println("No Token selected!")
		return;
	}
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	language="en";
	if m.Author.ID == s.State.User.ID {
		return
	}

	if (m.ChannelID == DE_CHANNEL){
		language = "de";
	}
	//CH Command
	if m.Content == "!ch" {
		commands.CMD_ch(s,m, language);
		RemoveMessage(s, m);
	}
	if m.Content == "!status" {
		commands.CMD_status(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!config" {
		commands.CMD_config(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!old" {
		commands.CMD_old(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!sp" {
		commands.CMD_sp(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!pro" {
		commands.CMD_pro(s,m, language);
		RemoveMessage(s, m);
	}

	if strings.Contains(m.Content, "!db") {
		commands.CMD_db(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!rechte" {
		commands.CMD_rechte(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!frage" {
		commands.CMD_frage(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!ksp" {
		commands.CMD_ksp(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!msv" {
		commands.CMD_msv(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!sf" {
		commands.CMD_sf(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!uv" {
		commands.CMD_uv(s,m, language);
		RemoveMessage(s, m);
	}

	if m.Content == "!uc" {
		commands.CMD_uc(s,m, language);
		RemoveMessage(s, m);
	}
	
	if m.Content == "!serverid" {
		commands.CMD_serverId(s,m, language);
		RemoveMessage(s, m);
	}
	
	if m.Content == "!passwordsec" {
		commands.CMD_passwordSec(s,m, language);
		RemoveMessage(s, m);
	}
	
	if m.Content == "!disnoti" {
		commands.CMD_discordNotify(s,m, language);
		RemoveMessage(s, m);
	}
	


	if m.Content == "!java" {
		s.ChannelMessageSend(m.ChannelID, "Was willst du mit Java? Ich bin ein GO Bot. Meine Meinung zu Java = 💩")
		RemoveMessage(s, m);
	}

	if m.Content == "!go" {
		s.ChannelMessageSend(m.ChannelID, "Jap, da gebe ich dir recht! GO ist der Boss im Haus... ")
		RemoveMessage(s, m);
	}

}

func RemoveMessage(s *discordgo.Session, m *discordgo.MessageCreate)  {
	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if (err != nil){
		fmt.Println("can't remove message")
	}
}
