package commands

import (
	"fmt"

	"gopkg.in/telegram-bot-api.v4"
)

type BatmanHandler struct {
}

var batmanHandlerInfo = CommandInfo{
	Command:     "batman",
	Args:        "",
	Permission:  3,
	Description: "says who is batman",
	LongDesc:    "",
	Usage:       "/batman",
	Examples: []string{
		"/batman",
	},
	ResType: "message",
}

func (h *BatmanHandler) HandleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, args []string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(GetUserTitle(message.From), " is Batman"))
	bot.Send(msg)
}

func (h *BatmanHandler) Info() *CommandInfo {
	return &batmanHandlerInfo
}

func (h *BatmanHandler) HandleReply(message *tgbotapi.Message) (bool, string) {
	return false, ""
}

func (h *BatmanHandler) Setup(setupFields map[string]interface{}) {

}
