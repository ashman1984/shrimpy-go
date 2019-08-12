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
