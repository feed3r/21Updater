package model

type Host struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type GHEventTranslator struct {
	Message     string `yaml:"message"`
	PushMessage string `yaml:"push_message"`
}

type Conf struct {
	Host       Host   `yaml:"host"`
	BotToken   string `yaml:"bot_token"`
	ChatId     string `yaml:"chat_id"`
	Logfile    string `yaml:"logfile"`
	Language   string `yaml:"lang"`
	Translator GHEventTranslator
}
