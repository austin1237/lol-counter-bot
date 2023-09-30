# lol-counter-bot
A discord bot that display possible counters for picks in leauge of legends [lol-counter-source-api](https://github.com/austin1237/lol-counter-api) and exposes for faster consumption.

## Prerequisites
You must have the following installed/configured on your system for this to work correctly<br />
1. [Docker](https://www.docker.com/)
2. [Docker-Compose](https://docs.docker.com/compose/)

## Environment Variables
The following variables need to be set on your local/ci system.

### COUNTER_API_URL
Url of the deployed lol-counter-api 
### DISCORD_BOT_TOKEN
token that grants bot access to your discord server.