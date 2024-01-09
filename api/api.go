package api

import (
	"encoding/json"
	"net/http"
)

// Coin balance params
type CoinBalanceParams struct {
	username string
}

// Coin balance response
type CoinBalanceResponse struct {
	code int
	balance int64
}

// Error response
type Error struct {
	code int
	message string
}