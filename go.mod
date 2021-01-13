module github.com/Support-pp/discord-bot

go 1.12

require (
	github.com/Support-pp/discord-bot/commands v0.0.0-00010101000000-000000000000
	github.com/bwmarrin/discordgo v0.22.1
)

replace github.com/Support-pp/discord-bot/commands => ./commands
