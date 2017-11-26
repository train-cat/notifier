package model

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"
)

// List of state available
const (
	StateDelayed = "delayed"
	StateDeleted = "deleted"
)

type (
	// PubSubMessage receive from pub/sub
	PubSubMessage struct {
		Message struct {
			Data string `json:"data"`
		} `json:"message"`
	}

	// Issue include in pubsub message
	Issue struct {
		State       string `json:"state"`
		Code        string `json:"code"`
		Schedule    string `json:"schedule"`
		StationID   int    `json:"station_id"`
		StationName string `json:"station_name"`
	}
)

// GetPubSubMessageFromHTTPRequest return PubSubMessage or error if they can't decode request body
func GetPubSubMessageFromHTTPRequest(r *http.Request) (*PubSubMessage, error) {
	m := &PubSubMessage{}

	err := json.NewDecoder(r.Body).Decode(m)

	return m, err
}

// GetIssueFromHTTPRequest return Issue or error if they can't decode PubSubMessage
func GetIssueFromHTTPRequest(r *http.Request) (*Issue, error) {
	m, err := GetPubSubMessageFromHTTPRequest(r)

	if err != nil {
		return nil, err
	}

	return m.GetIssue()
}

// GetIssue decode base64 content of PubSubMessage
func (m PubSubMessage) GetIssue() (*Issue, error) {
	data, err := base64.StdEncoding.DecodeString(m.Message.Data)

	if err != nil {
		return nil, err
	}

	r := Issue{}

	err = json.Unmarshal(data, &r)

	return &r, err
}

// GetSchedule return schedule formatted
func (i *Issue) GetSchedule() string {
	t, err := time.Parse("02/01/2006 15:04 -0700", i.Schedule)

	if err != nil {
		return ""
	}

	return t.Format("15:04")
}
