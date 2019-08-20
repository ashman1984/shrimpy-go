# shrimpy-go

A GoLang client for accessing the Shrimpy Developers API.
Takes care of authentication and http requests so you can get the data you need quickly.

- **Currently there is no Asset Management endpoint support**
- **Currently there is no WebSocket support**

## Usage

Install package from Github: `go get github.com/ashman1984/shrimpy-go`

Import the package at the top of your project: `import shrimpyclient "shrimpy-go"`

Instantiate your client config in your func main():
  ```
	var config shrimpyclient.Config
	config.Endpoint = "https://dev-api.shrimpy.io"
	config.MasterAPIKey = "Insert Master API Key Here"
	config.MasterSecretKey = "Insert Master Secret Key Here"
	config.DebugMessages = false

	sc := shrimpyclient.NewClient(config)
  ```
  
  From here you will just call inline functions like so:
  ```
  allSupportedExchanges := sc.GetSupportedExchanges()
  getBinanceAssets := sc.GetExchangeAssets("binance")
  getBinanceTradingPairs := sc.GetExchangePairs("binance")
  etc...
  ```
  Check the return types to view all properties returned from each function call.
  
  ## All Supported Functions
  
  **PUBLIC ENDPOINTS FUNCTIONS**
  - GetSupportedExchanges()
  - GetExchangeAssets(exchangeName string)
  - GetExchangePairs(exchangeName string)
  
  **MARKET ENDPOINTS FUNCTIONS**
  - GetExchangeTickers(exchangeName string)
  - GetCandleStickData(exchangeName string, quoteTradingSymbol string, baseTradingSymbol string, interval string)
  - GetOrderBooks(sliceExchanges []string, limit, quoteSymbol, baseSymbol string)
  
  **USER ENDPOINT FUNCTIONS**
  - GetUserList()
  - GetSingleUserList(userID string)
  - CreateUser(userName string)
  - RenameUser(userID string, userName string)
  - EnableUser(userID string)
  - DisableUser(userID string)
  
  **USER API KEY ENDPOINT FUNCTIONS**
  - GetAPIKeys(userID string)
  - CreateAPIKeys(userID string)
  - DeleteAPIKeys(userID string, publicKey string)
  - GetAPIKeyPermissions(userID string, publicKey string)
  - SetAPIKeyPermissions(userID string, publicKey string, tradePermission bool, accountPermission bool)
  
  **ACCOUNTS ENDPOINT FUNCTIONS**
  - ListAccounts(userID string)
  - GetAccount(userID string, exchangeAccountID string)
  - LinkExchangeAccount(userID string, exchangeName string, publicKey string, privateKey string, passphrase string)
  - UnLinkExchangeAccount(userID string, exchangeID string)
  - GetWhitelistedIPs(userID string)
  
  **TRADING ENDPOINT FUNCTIONS**
  - CreateTrade(userID string, exchangeID string, fromSymbol string, toSymbol string, amount string, smartRouting bool, maxSpreadPercent string, maxSlippagePercent string)
  - GetTradeStatus(userID string, exchangeID string, tradeID string)
  - GetActiveTrades(userID string, exchangeID string)
  
  **BALANCE ENDPOINT FUNCTIONS**
  - GetBalance(userID string, exchangeID string)
  - GetTotalBalanceHistory(userID string, exchangeID string)
  
  **LIMIT ORDER ENDPOINT FUNCTIONS**
  - PlaceLimitOrder(userID string, exchangeID string, baseSymbol string, quoteSymbol string, quantity string, side string, timeInForce string, price string)
  - GetLimitOrderStatus(userID string, exchangeID string, orderID string)
  - ListOpenOrders(userID string, exchangeID string)
  - CancelLimitOrder(userID string, exchangeID string, orderID string)
