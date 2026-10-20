<!--
SPDX-FileCopyrightText: NONE
SPDX-License-Identifier: Unlicense
-->
# Monica
A small little bot that works on both Discord and IRC.

## Platforms
Each platform has it's own executable, though they both share most of the same code

### Discord (Disgo)
To run it, simply run:
```sh
go run ./cmd/monica-disgo
```

And compile it with:
```sh
go build -o monica-disgo ./cmd/monica-disgo
```

### IRC (Lightning)
> [!NOTE]
> Due to issues with Lightning, it is planned on being dropped in favor of another IRC bot library

The same as with Discord but with a slightly different path, run it with:
```sh
go run ./cmd/monica-lightning
```

And to compile it, run:
```sh
go build -o monica-lightning ./cmd/monica-lightning
```

## FAQ

#### Why is this bot configured entirely through environment variables?
Why not? They are `.env` files loaded from `/etc/monica.env`, `~/.config/monica/monica.env`, and `./.env` in that order so you can configure it like you would any other tool (just with a different syntax of course)!

#### What commands are available?
Currently, the available commands are: `about`, `ping`, `define`, `calc`, `converttz`, `fortune`, `rtd`, `uwuify`, `ship`, `insult`, `compliment`, `animegirl`, and `flirt`.

#### Why is it called Monica?
Whilst thinking of what to name it, I noticed a lot of my bots were named after Anime characters, and there is one person in a server I am in who is **obsessed** with Monica Norton from [The Silent Witch](https://wikipedia.org/wiki/Secrets_of_the_Silent_Witch), so I named it Monica!

#### Why are all of the commands text-only?
The purpose of this bot is to be highly extensible frontend-wise, as such, all of the commands are currently text-only because every chat platform worth using supports text!