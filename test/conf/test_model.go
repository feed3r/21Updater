package conf

import "github.com/feed3r/21Updater/src/model"

const BASE_CONF_YAML = `
host:
  address: localhost
  port: 8080
bot_api_key: aaa_bbb_ccc
`

var ExpectedBaseConf = model.Conf{
	Host: model.Host{
		Address: "localhost",
		Port:    8080,
	},
	BotApiKey: "aaa_bbb_ccc",
}
