package test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/feed3r/21Updater/src/engine/telegram"
	"github.com/feed3r/21Updater/src/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestSendMessageToTelegram(t *testing.T) {

	var config model.Conf
	var logger *logrus.Logger

	data, err := os.ReadFile("../conf/conf.yaml")
	if err != nil {
		logger.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		t.Fatalf("Error parsing YAML: %v", err)
	}

	res, err := telegram.SendTextToTelegramChat(config.BotToken, config.ChatId, "Test message", logger)
	require.Nil(t, err)

	require.Equal(t, true, res.Ok)
	require.Equal(t, true, res.Result.From.IsBot)
	botTokenId, err := strconv.ParseInt(strings.Split(config.BotToken, ":")[0], 10, 64)
	require.Nil(t, err)
	require.Equal(t, botTokenId, res.Result.From.ID)
	chatId, err := strconv.ParseInt(config.ChatId, 10, 64)
	require.Nil(t, err)
	require.Equal(t, chatId, res.Result.Chat.ID)
	require.Equal(t, "Test message", res.Result.Text)
}
