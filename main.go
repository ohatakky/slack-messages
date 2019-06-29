package main

import (
	"os"

	"github.com/nlopes/slack"
)

func pushSlack(msg string) error {
	token := os.Getenv("TOKEN")
	api := slack.New(token)
	attachment := slack.Attachment{
		Pretext: msg,
	}
	cannelID := os.Getenv("CHANNELID")
	title := os.Getenv("TITLE")
	_, _, err := api.PostMessage(cannelID, slack.MsgOptionText(title, false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		return err
	}

	return nil
}
