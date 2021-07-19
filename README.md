# twitch-bot

A Twitch.tv Chat Bot

## Build and Run Commands

---

- build
  -- $go build
- run
  -- $./twitch-bot.exe
- develop
  -- $air

## Twitch Location

---

the Bot can be found running on Twitch under the name `dustinbrink_bot` at the Twitch channel [twitch.tv/dustinbrink_bot](https://www.twitch.tv/dustinbrink_bot)

## Configuration

---

Configuration can be found in local file ./config.json
The format is JSON and should contain these properties

| Property    | Description                                                                                                                                         |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| IrcUri      | The Uri for the twitch service (in SSL) [current documentation](https://dev.twitch.tv/docs/irc/guide) says this should be "irc.chat.twitch.tv:6697" |
| Nickname    | The IRC nickname the bot will use, should be all lowercase and same as twitch id                                                                    |
| OauthToken  | Twitch.tv Open Auth Token for READ/EDIT CHAT, generate one [here](https://twitchapps.com/tmi/) (may need to be periodically updated)                |
| IrcChannel  | The channel name you would like the bot to join                                                                                                     |
| SwansonUri  | Uri path to the Ron Swanson API, that generates a Ron Swanson quote                                                                                 |
| SslCertPath | Path to local file containing the ssl public key                                                                                                    |
| SslKeyPath  | Path to local file containing the private ssl key                                                                                                   |

## Overview

---

Create an automated [Twitch](https://dev.twitch.tv/docs/irc) chat bot console application that can be run from a command line interface (CLI).

## Requirements

---

The bot application should be able to:

- Console output all interactions - legibly formatted, with timestamps.
- Connect to Twitch IRC over SSL.
- Join a channel.
- Read a channel.
- Read a private message.
- Write to a channel
- Reply to a private message.
- Avoid premature disconnections by handling Twitch courier ping / pong requests.
- Publicly reply to a user-issued string command within a channel (!YOUR_COMMAND_NAME).
  - Reply to the "!swanson" command by dynamically returning a random Ron Swanson quote using the [Ron Swanson API](https://github.com/jamesseanwright/ron-swanson-quotes).

## Caveats

---

- The application must be written in Go using the [standard library](https://golang.org/pkg/) - absolutely no third-party module dependencies.
- All interactions should be asynchronous.
- The application should account for Twitch API rate limits.
- The application should not exit prematurely.
