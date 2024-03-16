package conf

import "github.com/feed3r/21Updater/src/model"

const BASE_CONF_YAML = `
host:
  address: localhost
  port: 8080
bot_api_key: aaa_bbb_ccc
chat_id: "12345"
logfile: "21updater.log"
`

var ExpectedBaseConf = model.Conf{
	Host: model.Host{
		Address: "localhost",
		Port:    8080,
	},
	BotToken: "aaa_bbb_ccc",
	ChatId:   "12345",
	Logfile:  "21updater.log",
}
