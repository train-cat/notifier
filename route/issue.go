package route

import (
	"net/http"

	"github.com/train-cat/notifier/api"
	"github.com/train-cat/notifier/bot"
	"github.com/train-cat/notifier/helper"
	"github.com/train-cat/notifier/model"
)

// Issue record an issue to the database
func Issue(w http.ResponseWriter, r *http.Request) {
	i, err := model.GetIssueFromHTTPRequest(r)

	if helper.HTTPError(w, err) {
		return
	}

	as, err := api.GetAlerts(i.Code, i.StationID)

	if helper.HTTPError(w, err) {
		return
	}

	if len(as) > 0 {
		err = bot.Notify(i, as)
	}

	if helper.HTTPError(w, err) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
