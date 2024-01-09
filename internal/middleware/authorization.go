package middleware

import (
	"errors"
	"net/http"

	"github.com/LawEyez/go-api/api/internal/tools"
	"github.com/LawEyez/go-api/api"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get username and token.
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		// Ensure valid username and token.
		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		// Get database.
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		// Get user from database.
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}