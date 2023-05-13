package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

type Status string

const (
	Success Status = "SUCCESS"
	Error   Status = "ERROR"
	Warning Status = "WARNING"
)

type SlackRequestBody struct {
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Color string `json:"color"`
	Text  string `json:"text"`
}

func SendNotificationToSlack(
	status Status,
	text string,
) error {
	var attachment Attachment
	switch status {
	case Success:
		attachment = Attachment{Color: "good", Text: text}
	case Error:
		attachment = Attachment{Color: "danger", Text: text}
	case Warning:
		attachment = Attachment{Color: "warning", Text: text}
	default:
		return errors.New("invalid status")
	}

	slackBody, _ := json.Marshal(SlackRequestBody{Attachments: []Attachment{attachment}})
	webhookUrl := os.Getenv("SLACK_WEBHOOK_URL")

	if webhookUrl == "" {
		log.Default().Println("webhook url  is not set")
		return nil
	}

	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("non-ok response returned from Slack")
	}
	return nil
}
