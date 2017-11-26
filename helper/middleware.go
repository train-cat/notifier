package helper

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var token string

// InitHelper init token var
func InitHelper() {
	token = viper.GetString("pubsub.token")
}

// MiddlewareSecurity is helper to ensure request sent contain the good token
func MiddlewareSecurity(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.URL.Query().Get("token")

		if t != token {
			log.Warningf("invalid token '%s'", t)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next(w, r)
	})
}

// MiddlewareNeedPost force request to be Post
func MiddlewareNeedPost(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Warning("Not a post request")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		next(w, r)
	})
}
