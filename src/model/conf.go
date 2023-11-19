package model

type Host struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}
type Conf struct {
	Host      Host   `yaml:"host"`
	BotApiKey string `yaml:"bot_api_key"`
}
