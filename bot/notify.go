package bot

import (
	log "github.com/sirupsen/logrus"
	"github.com/train-cat/client-train-go"
	"github.com/train-cat/notifier/model"
)

// Notify bot to send issue
func Notify(i *model.Issue, alerts []traincat.Alert) {
	n := notification{
		Issue: issue{
			Station:  i.StationName,
			Schedule: i.GetSchedule(),
			State:    i.State,
		},
		Actions: []action{},
	}

	for _, alert := range alerts {
		a := traincat.Action{}

		err := alert.Embedded.Get("action", &a)

		if err != nil {
			log.Errorf("[notify] %s", err)
			continue
		}

		n.Actions = append(n.Actions, action{
			Type: a.Type,
			Data: a.Data,
		})
	}

	err := n.send()

	if err != nil {
		log.Errorf("[send notification] %s", err)
	}
}
