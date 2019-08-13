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
	DebugMessages   bool
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

//WhitelistedIPs returns all whitelisted IPS associated with this account
type WhitelistedIPs []string

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

//LinkedAccounts for storing linked account data
type LinkedAccounts []struct {
	ID                int           `json:"id"`
	Exchange          string        `json:"exchange"`
	IsRebalancing     bool          `json:"isRebalancing"`
	ExchangeAPIErrors []interface{} `json:"exchangeApiErrors"`
}

//LinkedExchangeAccount for storing exchange account data
type LinkedExchangeAccount struct {
	ID                int           `json:"id"`
	Exchange          string        `json:"exchange"`
	IsRebalancing     bool          `json:"isRebalancing"`
	ExchangeAPIErrors []interface{} `json:"exchangeApiErrors"`
}

//LinkAccountRequest is the data we send to link an account to a user
type LinkAccountRequest struct {
	Exchange   string `json:"exchange"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	Passphrase string `json:"passphrase,omitempty"`
}

//LinkAccountResponse is the responding id for the account that was just linked
type LinkAccountResponse struct {
	ID int `json:"id"`
}

//CreateTradeResponse stores the trade id of a posted trade
type CreateTradeResponse struct {
	ID string `json:"id"`
}

//CreateTradeRequest is the body for a create trade request
type CreateTradeRequest struct {
	FromSymbol         string `json:"fromSymbol"`
	ToSymbol           string `json:"toSymbol"`
	Amount             string `json:"amount"`
	SmartRouting       bool   `json:"smartRouting,omitempty"`
	MaxSpreadPercent   string `json:"maxSpreadPercent,omitempty"`
	MaxSlippagePercent string `json:"maxSlippagePercent,omitempty"`
}

//TradeStatus holds info from requesting status of a particular trade
type TradeStatus struct {
	Trade struct {
		ID                   string        `json:"id"`
		FromSymbol           string        `json:"fromSymbol"`
		ToSymbol             string        `json:"toSymbol"`
		Amount               string        `json:"amount"`
		Status               string        `json:"status"`
		Success              bool          `json:"success"`
		ErrorCode            int           `json:"errorCode"`
		ErrorMessage         string        `json:"errorMessage"`
		ExchangeAPIErrors    []interface{} `json:"exchangeApiErrors"`
		SmartRouting         bool          `json:"smartRouting"`
		MaxSpreadPercent     string        `json:"maxSpreadPercent"`
		MaxSlippagePercent   string        `json:"maxSlippagePercent"`
		TriggeredMaxSpread   bool          `json:"triggeredMaxSpread"`
		TriggeredMaxSlippage bool          `json:"triggeredMaxSlippage"`
	} `json:"trade"`
	Changes []struct {
		Symbol      string  `json:"symbol"`
		NativeValue string  `json:"nativeValue"`
		BtcValue    float64 `json:"btcValue"`
		UsdValue    float64 `json:"usdValue"`
	} `json:"changes"`
	Fills []struct {
		BaseAmount  string  `json:"baseAmount"`
		BaseSymbol  string  `json:"baseSymbol"`
		BtcValue    float64 `json:"btcValue"`
		Price       string  `json:"price"`
		QuoteAmount string  `json:"quoteAmount"`
		QuoteSymbol string  `json:"quoteSymbol"`
		Side        string  `json:"side"`
		UsdValue    float64 `json:"usdValue"`
	} `json:"fills"`
}

//ActiveTrades holds an array of active trades associated with a useraccount / exchange
type ActiveTrades []struct {
	ID                   string        `json:"id"`
	FromSymbol           string        `json:"fromSymbol"`
	ToSymbol             string        `json:"toSymbol"`
	Amount               float64       `json:"amount"`
	Status               string        `json:"status"`
	Success              bool          `json:"success"`
	ErrorCode            int           `json:"errorCode"`
	ErrorMessage         string        `json:"errorMessage"`
	ExchangeAPIErrors    []interface{} `json:"exchangeApiErrors"`
	SmartRouting         bool          `json:"smartRouting"`
	MaxSpreadPercent     string        `json:"maxSpreadPercent"`
	MaxSlippagePercent   string        `json:"maxSlippagePercent"`
	TriggeredMaxSpread   bool          `json:"triggeredMaxSpread"`
	TriggeredMaxSlippage bool          `json:"triggeredMaxSlippage"`
}

//TotalBalanceHistory holds aggregate balance data
type TotalBalanceHistory []struct {
	Date     time.Time `json:"date"`
	UsdValue float64   `json:"usdValue"`
	BtcValue float64   `json:"btcValue"`
}

//LimitOrderRequest holds data for a limit order request
type LimitOrderRequest struct {
	BaseSymbol  string `json:"baseSymbol"`
	QuoteSymbol string `json:"quoteSymbol"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Side        string `json:"side"`
	TimeInForce string `json:"timeInForce"`
}

//LimitOrderReturn returns ID of order
type LimitOrderReturn struct {
	ID string `json:"id"`
}

//LimitOrderStatusReturn holds a particular executed orders information
type LimitOrderStatusReturn struct {
	Order struct {
		ID                string        `json:"id"`
		BaseSymbol        string        `json:"baseSymbol"`
		QuoteSymbol       string        `json:"quoteSymbol"`
		Amount            string        `json:"amount"`
		Price             string        `json:"price"`
		Side              string        `json:"side"`
		TimeInForce       string        `json:"timeInForce"`
		Status            string        `json:"status"`
		CancelRequested   bool          `json:"cancelRequested"`
		Success           bool          `json:"success"`
		ErrorCode         int           `json:"errorCode"`
		ErrorMessage      string        `json:"errorMessage"`
		ExchangeAPIErrors []interface{} `json:"exchangeApiErrors"`
	} `json:"order"`
	Changes []struct {
		Symbol      string  `json:"symbol"`
		NativeValue string  `json:"nativeValue"`
		BtcValue    float64 `json:"btcValue"`
		UsdValue    float64 `json:"usdValue"`
	} `json:"changes"`
}

//OpenActiveOrders holds all not in 'completed' state orders
type OpenActiveOrders []struct {
	Amount            string        `json:"amount"`
	Price             string        `json:"price"`
	CancelRequested   bool          `json:"cancelRequested"`
	ErrorCode         int           `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
	BaseSymbol        string        `json:"baseSymbol"`
	QuoteSymbol       string        `json:"quoteSymbol"`
	ID                string        `json:"id"`
	Status            string        `json:"status"`
	Success           bool          `json:"success"`
	Side              string        `json:"side"`
	TimeInForce       string        `json:"timeInForce"`
	ExchangeAPIErrors []interface{} `json:"exchangeApiErrors"`
}
