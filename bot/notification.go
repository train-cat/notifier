package bot

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

var client = http.Client{}

type (
	notification struct {
		Issue   issue    `json:"issue"`
		Actions []action `json:"actions"`
	}

	issue struct {
		Station  string `json:"station"`
		Schedule string `json:"schedule"`
		State    string `json:"state"`
	}

	action struct {
		Data map[string]string `json:"data"`
		Type string            `json:"type"`
	}
)

func (n *notification) send() error {
	url := viper.GetString("bot.webhook")

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	if err := enc.Encode(n); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(buf.Bytes()))

	if err != nil {
		return nil
	}

	_, err = client.Do(req)

	return err
}
