# Discord Go Adventures
A fun tabletop RPG bot built for Game Masters. This is a work-in-progress, and the first system I'm building out will be the Mistborn Adventure Game by Crafty.

__To Do:__
- [ ] character sheet and actions loaded using YAML or TOML
- [ ] generate rolls from your sheet instead of hardcoded declarations
- [ ] on roll stop/end update display (embedded message) with the complete log of actions and results instead of deleting the message
- [ ] interfaces for adding games other than mistborn

## Getting Started
You will need a bot token, you can get one here: https://discord.com/developers. This bot will need read/write access to:
- the server you want it to participate in
- the channels you want it to interact in
- the roles (api calls these guilds) you will want it to be aware of (pending work will take advantage of this)
```shell
$ go run cmd/dgadventures/main.go -t <your_discord_token_here>
```

## Credit

### Libraries
- [discordgo](https://github.com/bwmarrin/discordgo)
- [dgrouter](https://github.com/Necroforger/dgrouter)

### Inspiration
- [dgwidgets](https://github.com/Necroforger/dgwidgets)

