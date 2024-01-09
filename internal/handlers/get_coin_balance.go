package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LawEyez/go-api/api"
	"github.com/LawEyez/go-api/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinsBalance(w http.ResponseWriter, r *http.Request) {
	// Get request params.
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Get database.
	database, err := tools.NewDatabase()

	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Get coins.
	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)

	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Set value to our response.
	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}