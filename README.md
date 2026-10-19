<!--
SPDX-FileCopyrightText: NONE
SPDX-License-Identifier: Unlicense
-->
# Monica
A small little bot that works on Discord, and IRC

To run it, simply run:
```sh
go run ./cmd
```

To compile it, simply run:
```sh
go build -o monica ./cmd
```

## FAQ

#### Why is this bot configured entirely through environment variables?
Why not? They are `.env` files loaded from `/etc/monica.env`, `~/.config/monica/monica.env`, and `./.env` in that order so you can configure it like you would any other tool (just with a different syntax of course)!

#### What commands are available?
Currently, the available commands are: `about`, `ping`, `define`, `calc`, `fortune`, `uwuify`, `ship`, `insult`, `compliment`, `animegirl`, and `flirt`.

#### Why is it called Monica?
Whilst thinking of what to name it, I noticed a lot of my bots were named after Anime characters, and there is one person in a server I am in who is **obsessed** with Monica Norton from [The Silent Witch](https://wikipedia.org/wiki/Secrets_of_the_Silent_Witch), so I named it Monica!
