Test commit


# TelegramBOT
A simple program, written in GO, that receives events from a GitHub repository and sends message to a Telegram chat using a Telegram BOT. 

Actually, this program creates a web server on a configured port to receive the events. Once the event is received, a proper message is created, and it's sent to a specific Telegram Chat through a specific Telegram Bot. 

# Preliminary informations
This program receives events from a GitHub repository, using WebHooks. These events can be configured in the **WebHooks** section of the **Settings** page of the repository. Actually this program manages only PULL REQUESTS and ISSUE events. 

For Telegram, some preliminary data are mandatory to get the system to work, first of all a new Telegram Bot should be created, this can be easily achieved through the BotFather Telegram chat (see https://core.telegram.org/bots/tutorial)

Then, a ChatID should be provided. This represents the identifier of a specific Telegram Chat, to get the ChatID you can: 

* Simply go to IDBot chat on Telegram and write /getid command. This will return your personal chat ID, this can be useful to make some tests and receive the updates in your personal account.

* To get the chatID of a Group Chat, the procedure is a bit more complex, and it requires that you create the Bot (via the BotFather chat) and add it to the group chat you want to send the message. Once done, can interrogate the following link:

```
https://api.telegram.org/bot<YourBOTToken>/getUpdates
```

for example via curl:

```shell
curl https://api.telegram.org/bot<YourBOTToken>/getUpdates
```

this will return some informations, where you'll find also:

```json
{
    "update_id": 8393,
    "message": {
        "message_id": 3,
        "from": {
            "id": 7474,
            "first_name": "AAA"
        },
        "chat": {
            "id": "<Group ID>,"
            "title": "<Group name>"
        },
        "date": 25497,
        "new_chat_participant": {
            "id": 71, 
            "first_name": "NAME",
            "username": "YOUR_BOT_NAME"
        }
    }
}
```

here the "chat" struct can be found and the "id" attribute is the information you are looking for.


# Program usage
From the project main directory, the following command can be used to compile the program and produce an executable file called 21Updater:

```shell
go build -o 21Updater src/main.go
```

and then the program can be called using: 

```shell
./21Updater
```

Parameter --help can be used to retrieve the possible command line arguments:

```shell
./21Updater --help

Usage of ./main:
  -conf string
    	path to the configuration file (default "./conf/conf.yaml")
```

An example of the configuration file is provided below:
```yaml
host:     
  address: 'localhost'
  port: 8080
bot_token: 'YOUR_BOT_TELEGRAM_TOKEN_HERE'
chat_id: 'YOUR_CHAT_ID_HERE'
```

where:

**host**: represents a struct describing the host where the server will run. The two sub parameters are:
  **address**: address of the server host (typically: localhost). This can also be left empty to allow server to receive connection from any network interface.

  **port**: specify the server port where the server will listen (this is mandatory)

**bot_token**: specify the secret bot token obtained by Telegram (this should be taken by The BotFather, see https://core.telegram.org/bots/tutorial).

**chat_id**: specify the ID of the target chat where the bot should write. 

Once the configuration is done, you can simply run the program:

```
./21Updater
```

producing the following message:

```
21Updater server started and listening at :8080
```

Now, once an event is received from the GitHub repository, the received message is printed on the standard output.






