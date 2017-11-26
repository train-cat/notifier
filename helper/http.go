package helper

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HTTPError log error and write 500 status code to http.ResponseWriter and return true if err is not nil
func HTTPError(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return true
	}

	return false
}
