package bot

import (
	"github.com/train-cat/client-train-go"
	"github.com/train-cat/notifier/model"
)

// Notify bot to send issue
func Notify(i *model.Issue, alerts []traincat.Alert) error {
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
			return err
		}

		n.Actions = append(n.Actions, action{
			Type: a.Type,
			Data: a.Data,
		})
	}

	return n.send()
}
