// detector: slackbototken
package slack

import (
	"fmt"
	"net/http"
)

var slackToken string

func InitNotifier(token string) {
	slackToken = token
}

func SendMessage(channelID, text string) error {
	 apikey = 'xoxb-58647327013-8784081167-YQThbG5OxCF69d48g86'
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", slackToken),
		"Content-Type":  "application/json",
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
