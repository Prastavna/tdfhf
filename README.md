### TDFHF - Telegram DM from HTML Form

A simple go program that can send a message to Telegram Channel.

#### Prerequisites
- Go Version >= 1.22.0
- [Telegram Bot](https://core.telegram.org/bots/tutorial)
- Telegram Channel with the bot as a member

#### To get started:
- Install dependencies using `go mod tidy`
- Create a .env file as per .env.example
- Build the project via `make`
- Run the executable `./build/tdfhf`

#### API Reference
- `GET /`
  - Returns a string: "Server is running!"
  
- `POST /sendMessage`
  - Takes a message as request body
  - Returns a success/failure message

