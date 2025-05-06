# Chatty Charles 🤖

```
              )\       _
     .--._  ,'  \_.-~~/'
     \    \'_    __ ( _ _
       \   (_)  |__|/~ ~~=~\
         )_____.---~~ \>\~-./'
        /'    //===  /==(
       (     /' __\ ( __\)
       ( /~\(   o   |_o_(
       (( (         _____)
        \\_/     ,      )
         \ \   ~-.._  /
        ._)/ \        /
         \/   \     ./
          /     \~/~'
        /       /
       '~~-.__./
```

A friendly CLI chatbot powered by OpenAI's GPT models.

## Features

- Interactive command-line interface
- Maintains conversation history
- Uses OpenAI's GPT models (gpt-3.5-turbo by default)
- Colorful ASCII art welcome and goodbye messages
- Graceful exit handling

## Prerequisites

- Go 1.24 or later
- OpenAI API key

## Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/chatty-charles.git
cd chatty-charles
```

2. Install dependencies:

```bash
go mod download
```

3. Set up your OpenAI API key:

```bash
# Create a .env file
echo "OPENAI_API_KEY=your-api-key-here" > .env
```

## Running the Chatbot

1. Navigate to the chatbot directory:

```bash
cd chatbot
```

2. Run the chatbot:

```bash
go run .
```

3. Start chatting! Type 'exit' to quit.

## Configuration

You can configure the chatbot by setting environment variables:

- `OPENAI_API_KEY`: Your OpenAI API key (required)
- `MODEL`: The GPT model to use (default: gpt-3.5-turbo)
- `TEMPERATURE`: Response randomness (default: 0.7)

## Project Structure

```
chatbot/
├── main.go      # CLI entry point
├── config.go    # Configuration handling
├── client.go    # OpenAI API client
├── chat.go      # Conversation handling
└── display.go   # ASCII art and display functions
```

## License

MIT License
