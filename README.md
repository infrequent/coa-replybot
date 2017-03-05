#coa-replybot


This is a chatbot example based on the DiscordGo, PingPong Example Bot.
This Bot will respond to 
```Hello```
with
```:wave: Welcome to the Government Service! How can I help you today?``` 
and
```I'd like to find information about paying rates, please!```
with 
```Sure, here is the link you're after: http://govt.org.au/rates``` 

<img  src="http://i.imgur.com/XkKye91.png">
## Build

This assumes you already have a working Go environment setup and that
DiscordGo is correctly installed on your system.

```sh
go build
```

####Help with installing Go environment setup and DiscordGo
* Installing DiscordGO - [https://github.com/bwmarrin/discordgo#installing](https://github.com/bwmarrin/discordgo#installing)
* Installing GoLang - [https://github.com/golang/go/wiki/Ubuntu#install-package-golang](https://github.com/golang/go/wiki/Ubuntu#install-package-golang)
* Setting GOPATH example in "~/go": ```echo 'export GOPATH=$HOME/go' >> ~/.bashrc```



## Usage

This example uses bot tokens for authentication only.
While user/password is supported by DiscordGo, it is not recommended.

```
./coa-replybot --help
Usage of ./coa-replybot:
  -t string
        Bot Token
```

The below example shows how to start the bot

```sh
./coa-replybot -t YOUR_BOT_TOKEN
```

Creating a discord bot account and retrieving a token - [https://github.com/reactiflux/discord-irc/wiki/Creating-a-discord-bot-&-getting-a-token](https://github.com/reactiflux/discord-irc/wiki/Creating-a-discord-bot-&-getting-a-token)
## Attributions

* Airhorn - [https://github.com/hammerandchisel/airhornbot](https://github.com/hammerandchisel/airhornbot)
* DiscordGo - [https://github.com/bwmarrin/discordgo](https://github.com/bwmarrin/discordgo)
* DiscordGo Examples - [https://github.com/bwmarrin/discordgo/tree/master/examples](https://github.com/bwmarrin/discordgo/tree/master/examples)

## Copyright and License
DiscordGo License - [https://github.com/bwmarrin/discordgo/blob/master/LICENSE](https://github.com/bwmarrin/discordgo/blob/master/LICENSE)