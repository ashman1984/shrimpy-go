package shrimpygo

import (
	"time"
)

//Client is the connection
type Client struct {
	Config Config
}

//Config for the client to work
type Config struct {
	Endpoint        string
	MasterAPIKey    string
	MasterSecretKey string
	TimeoutSecond   int
}

//SupportedExchanges gets all exchanges that shrimpy suppports
type SupportedExchanges []struct {
	Exchange     string  `json:"exchange"`
	BestCaseFee  float64 `json:"bestCaseFee"`
	WorstCaseFee float64 `json:"worstCaseFee"`
	Icon         string  `json:"icon"`
}

//Assets are coin objects for an exchange
type Assets []struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	TradingSymbol string `json:"tradingSymbol"`
}

//Pairs are coin pairs for an exchange
type Pairs []struct {
	BaseTradingSymbol  string `json:"baseTradingSymbol"`
	QuoteTradingSymbol string `json:"quoteTradingSymbol"`
}

//Tickers are tickers for an exchange
type Tickers []struct {
	Name                string    `json:"name"`
	Symbol              string    `json:"symbol"`
	PriceUsd            string    `json:"priceUsd"`
	PriceBtc            string    `json:"priceBtc"`
	PercentChange24HUsd string    `json:"percentChange24hUsd"`
	LastUpdated         time.Time `json:"lastUpdated"`
}

//ExchangeOrders - list of exchange objects which contains orderbooks
type ExchangeOrders []struct {
	QuoteSymbol string `json:"quoteSymbol"`
	BaseSymbol  string `json:"baseSymbol"`
	OrderBooks  []struct {
		Exchange  string `json:"exchange"`
		OrderBook struct {
			Asks []struct {
				Price    string `json:"price"`
				Quantity string `json:"quantity"`
			} `json:"asks"`
			Bids []struct {
				Price    string `json:"price"`
				Quantity string `json:"quantity"`
			} `json:"bids"`
		} `json:"orderBook"`
	} `json:"orderBooks"`
}

//CandleSticks hold all candle sticks for requested period
type CandleSticks []struct {
	Open        string    `json:"open"`
	High        string    `json:"high"`
	Low         string    `json:"low"`
	Close       string    `json:"close"`
	Volume      string    `json:"volume"`
	QuoteVolume float64   `json:"quoteVolume"`
	BtcVolume   float64   `json:"btcVolume"`
	UsdVolume   float64   `json:"usdVolume"`
	Time        time.Time `json:"time"`
}

//UsersList defines a list of users associated with this masterAPI key
type UsersList []struct {
	ExpirationDate time.Time `json:"expirationDate"`
	ID             string    `json:"id"`
	IsEnabled      bool      `json:"isEnabled"`
	Name           string    `json:"name"`
}

//SingleUser defines a single user object
type SingleUser struct {
	ExpirationDate time.Time `json:"expirationDate"`
	ID             string    `json:"id"`
	IsEnabled      bool      `json:"isEnabled"`
	Name           string    `json:"name"`
}

//UserID of new user
type UserID struct {
	ID string `json:"id"`
}

//CreateUserRequest optional field creating new user
type CreateUserRequest struct {
	Name string `json:"name"`
}

//SuccessReturn returns whether the operation was successful
type SuccessReturn struct {
	Success bool `json:"success"`
}

//GetPublicAPIKeys returns all public keys associated with this user
type GetPublicAPIKeys []string

//CreateAPIKeyReturn creates new keys for this user
type CreateAPIKeyReturn struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

//APIKeyPermissions holds shrimpy permissions for a particular public key
type APIKeyPermissions struct {
	Account bool `json:"account"`
	Trade   bool `json:"trade"`
}
