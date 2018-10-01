package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"spGoBot/commands"
	"strings"
)



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

	if m.Author.ID == s.State.User.ID {
		return
	}

	//CH Command
	if m.Content == "!ch" {
		commands.CMD_ch(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!config" {
		commands.CMD_config(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!old" {
		commands.CMD_old(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!sp" {
		commands.CMD_sp(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!pro" {
		commands.CMD_pro(s,m);
		RemoveMessage(s, m);
	}

	if strings.Contains(m.Content, "!db") {
		commands.CMD_db(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!rechte" {
		commands.CMD_rechte(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!frage" {
		commands.CMD_frage(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!ksp" {
		commands.CMD_ksp(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!sf" {
		commands.CMD_sf(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!uv" {
		commands.CMD_uv(s,m);
		RemoveMessage(s, m);
	}

	if m.Content == "!uc" {
		commands.CMD_uc(s,m);
		RemoveMessage(s, m);
	}


	if m.Content == "!java" {
		s.ChannelMessageSend(m.ChannelID, "Was willst du mit Java? Ich bin ein GO Bot. Meine Meinung zu Java = 💩")
	}

	if m.Content == "!go" {
		s.ChannelMessageSend(m.ChannelID, "Jap, da gebe ich dir recht! GO ist der Boss im Haus... ")
	}

}

func RemoveMessage(s *discordgo.Session, m *discordgo.MessageCreate)  {
	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if (err != nil){
		fmt.Println("can't remove message")
	}
}
