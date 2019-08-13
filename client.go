package shrimpygo

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	//GET so I don't have to type quotes all the time
	GET = "GET"
	//POST so I don't have to type quotes all the time
	POST = "POST"
	//DELETE so I don't have to type quotes all the time
	DELETE = "DELETE"
)

var nonce int64

/*

	START PUBLIC FUNCTIONS

*/

//GetSupportedExchanges returns a list of exchange objects for all exchanges that shrimpy supports
func (client *Client) GetSupportedExchanges() SupportedExchanges {
	r := new(SupportedExchanges)
	jsonStringReturn := httpDo(GET, "", "/v1/list_exchanges", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetExchangeAssets returns a list of exchange assets for a particular exchange
func (client *Client) GetExchangeAssets(exchangeName string) Assets {
	r := new(Assets)
	jsonStringReturn := httpDo(GET, "", "/v1/exchanges/"+exchangeName+"/assets", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetExchangePairs returns a list of exchange pairs for a particular exchange
func (client *Client) GetExchangePairs(exchangeName string) Pairs {
	r := new(Pairs)
	jsonStringReturn := httpDo(GET, "", "/v1/exchanges/"+exchangeName+"/trading_pairs", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

/*

	END PUBLIC FUNCTIONS
	START MARKET FUNCTIONS

*/

//GetExchangeTickers returns a list of exchange tickers for all exchanges that shrimpy supports
func (client *Client) GetExchangeTickers(exchangeName string) Tickers {
	r := new(Tickers)
	jsonStringReturn := httpDo(GET, "", "/v1/exchanges/"+exchangeName+"/ticker", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetCandleStickData returns all candlesticks for this exchange
func (client *Client) GetCandleStickData(exchangeName string, quoteTradingSymbol string, baseTradingSymbol string, interval string) CandleSticks {
	r := new(CandleSticks)

	params := "?quoteTradingSymbol=" + quoteTradingSymbol + "&baseTradingSymbol=" + baseTradingSymbol + "&interval=" + interval

	jsonStringReturn := httpDo(GET, params, "/v1/exchanges/"+exchangeName+"/candles", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetOrderBooks returns orderbooks for all exchanges in the slice, basesymbol, quotesymbol and limit are all inputs
func (client *Client) GetOrderBooks(sliceExchanges []string, limit, quoteSymbol, baseSymbol string) ExchangeOrders {

	r := new(ExchangeOrders)

	exchange := "all"
	sLimit := "&limit=10"
	qSymbol := ""
	bSymbol := ""

	if len(sliceExchanges) > 0 {
		exchange = strings.Join(sliceExchanges, ",")
	}

	if limit != "" {
		sLimit = "&limit=" + limit
	}

	if quoteSymbol != "" {
		qSymbol = "&quoteSymbol=" + quoteSymbol
	}

	if baseSymbol != "" {
		bSymbol = "&baseSymbol=" + baseSymbol
	}

	params := "?exchange=" + exchange + sLimit + qSymbol + bSymbol

	jsonStringReturn := httpDo(GET, params, "/v1/orderbooks", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

/*

	END MARKET FUNCTIONS
	START USER FUNCTIONS - NOTE: You must have the user permission enabled on your Master API key to use these endpoints

*/

//GetUserList returns all users associated with your masterAPI key
func (client *Client) GetUserList() UsersList {
	r := new(UsersList)

	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetSingleUserList returns a single user based on ID
func (client *Client) GetSingleUserList(userID string) SingleUser {
	r := new(SingleUser)

	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID, "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//CreateUser associates a new user with your account. Returns the ID of that user
func (client *Client) CreateUser(userName string) UserID {
	r := new(UserID)
	params := ""
	finalBody := ""

	if userName != "" {
		var body CreateUserRequest
		body.Name = userName

		stringBody, err := json.Marshal(body)

		if err != nil {
			fmt.Println(err)
		}

		finalBody = string(stringBody)
	}

	jsonStringReturn := httpDo(POST, params, "/v1/users", finalBody, client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//RenameUser names the user associated with the userID to the string userName
func (client *Client) RenameUser(userID string, userName string) SuccessReturn {
	r := new(SuccessReturn)
	params := ""

	var body CreateUserRequest
	body.Name = userName

	stringBody, err := json.Marshal(body)

	if err != nil {
		fmt.Println(err)
	}

	finalBody := string(stringBody)

	jsonStringReturn := httpDo(POST, params, "/v1/users/"+userID+"/name", finalBody, client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//EnableUser enables the user with the given ID
func (client *Client) EnableUser(userID string) SuccessReturn {
	r := new(SuccessReturn)
	params := ""

	jsonStringReturn := httpDo(POST, params, "/v1/users/"+userID+"/enable", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//DisableUser enables the user with the given ID
func (client *Client) DisableUser(userID string) SuccessReturn {
	r := new(SuccessReturn)
	params := ""

	jsonStringReturn := httpDo(POST, params, "/v1/users/"+userID+"/disable", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

/*

	END USER FUNCTIONS
	START USER API KEY FUNCTIONS

*/

//GetAPIKeys gets the public keys associated with this user
func (client *Client) GetAPIKeys(userID string) GetPublicAPIKeys {
	r := new(GetPublicAPIKeys)
	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID+"/keys", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//CreateAPIKeys creates a new public and private key for this user
func (client *Client) CreateAPIKeys(userID string) CreateAPIKeyReturn {
	r := new(CreateAPIKeyReturn)
	params := ""

	jsonStringReturn := httpDo(POST, params, "/v1/users/"+userID+"/keys", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//DeleteAPIKeys deletes a public key
func (client *Client) DeleteAPIKeys(userID string, publicKey string) SuccessReturn {
	r := new(SuccessReturn)
	params := ""

	jsonStringReturn := httpDo(DELETE, params, "/v1/users/"+userID+"/keys/"+publicKey, "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetAPIKeyPermissions returns the shrimpy permissions for this key
func (client *Client) GetAPIKeyPermissions(userID string, publicKey string) APIKeyPermissions {
	r := new(APIKeyPermissions)
	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID+"/keys/"+publicKey+"/permissions", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//SetAPIKeyPermissions sets the shrimpy permissions for this key
func (client *Client) SetAPIKeyPermissions(userID string, publicKey string, tradePermission bool, accountPermission bool) SuccessReturn {
	r := new(SuccessReturn)
	params := ""

	var permissions APIKeyPermissions
	permissions.Account = accountPermission
	permissions.Trade = tradePermission
	stringBody, err := json.Marshal(permissions)

	if err != nil {
		fmt.Println(err)
	}

	finalBody := string(stringBody)
	jsonStringReturn := httpDo(POST, params, "/v1/users/"+userID+"/keys/"+publicKey+"/permissions", finalBody, client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

/*

	END USER API KEYS FUNCTIONS
	START ACCOUNTS FUNCTIONS

*/

//ListAccounts will return an array of exchange accounts linked with this useraccount
func (client *Client) ListAccounts(userID string) LinkedAccounts {
	r := new(LinkedAccounts)
	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID+"/accounts", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetAccount will return a singular linked exchange object
func (client *Client) GetAccount(userID string, exchangeAccountID string) LinkedExchangeAccount {
	r := new(LinkedExchangeAccount)
	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID+"/accounts/"+exchangeAccountID, "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//LinkExchangeAccount will link your exchange keys with this user account
func (client *Client) LinkExchangeAccount(userID string, exchangeName string, publicKey string, privateKey string, passphrase string) LinkAccountResponse {
	r := new(LinkAccountResponse)
	params := ""

	var body LinkAccountRequest
	body.Exchange = exchangeName
	body.PrivateKey = privateKey
	body.PublicKey = publicKey

	if passphrase != "" {
		body.Passphrase = passphrase
	}

	stringBody, err := json.Marshal(body)

	if err != nil {
		fmt.Println(err)
	}

	finalBody := string(stringBody)
	fmt.Println(finalBody)

	jsonStringReturn := httpDo(POST, params, "/v1/users/"+userID+"/accounts", finalBody, client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//UnLinkExchangeAccount will unlike this exchange from your user account
func (client *Client) UnLinkExchangeAccount(userID string, exchangeID string) SuccessReturn {
	r := new(SuccessReturn)
	params := ""

	jsonStringReturn := httpDo(DELETE, params, "/v1/users/"+userID+"/accounts/"+exchangeID, "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetWhitelistedIPs will unlike this exchange from your user account
func (client *Client) GetWhitelistedIPs(userID string) WhitelistedIPs {
	r := new(WhitelistedIPs)
	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID+"/whitelist", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

/*

	END ACCOUNTS FUNCTIONS
	START TRADING FUNCTIONS

*/

//CreateTrade will post a trade for this user to this exchange
func (client *Client) CreateTrade(userID string, exchangeID string, fromSymbol string, toSymbol string, amount string, smartRouting bool, maxSpreadPercent string, maxSlippagePercent string) CreateTradeResponse {
	r := new(CreateTradeResponse)
	params := ""

	var body CreateTradeRequest
	body.Amount = amount
	body.FromSymbol = fromSymbol
	body.ToSymbol = toSymbol

	if smartRouting {
		body.SmartRouting = smartRouting
	}

	if maxSpreadPercent != "" {
		body.MaxSpreadPercent = maxSpreadPercent
	}

	if maxSlippagePercent != "" {
		body.MaxSlippagePercent = maxSlippagePercent
	}

	stringBody, err := json.Marshal(body)

	if err != nil {
		fmt.Println(err)
	}

	finalBody := string(stringBody)
	fmt.Println(finalBody)

	jsonStringReturn := httpDo(POST, params, "/v1/users/"+userID+"/accounts/"+exchangeID+"/trades", finalBody, client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetTradeStatus will return the details of a particular trade
func (client *Client) GetTradeStatus(userID string, exchangeID string, tradeID string) TradeStatus {
	r := new(TradeStatus)
	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID+"/accounts/"+exchangeID+"/trades/"+tradeID, "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetActiveTrades will return all trades who status is not 'completed'
func (client *Client) GetActiveTrades(userID string, exchangeID string) ActiveTrades {
	r := new(ActiveTrades)
	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID+"/accounts/"+exchangeID+"/trades", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

/*

	END TRADING FUNCTIONS
	START BALANCE FUNCTIONS

*/

//GetBalance will return the balances on all held assets on that exchange
func (client *Client) GetBalance(userID string, exchangeID string) ActiveTrades {
	r := new(ActiveTrades)
	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID+"/accounts/"+exchangeID+"/balance", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)

	if client.Config.DebugMessages {
		fmt.Println(jsonStringReturn)
	}

	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

//GetTotalBalanceHistory gets an aggregate balance history for an exchange account
func (client *Client) GetTotalBalanceHistory(userID string, exchangeID string) TotalBalanceHistory {
	r := new(TotalBalanceHistory)
	params := ""

	jsonStringReturn := httpDo(GET, params, "/v1/users/"+userID+"/accounts/"+exchangeID+"/total_balance_history", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)

	if client.Config.DebugMessages {
		fmt.Println(jsonStringReturn)
	}

	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
}

/*

	END BALANCE FUNCTIONS
	START LIMIT ORDER FUNCTIONS

*/

/*

	END LIMIT ORDER FUNCTIONS

*/

//NewClient initiates a new client object
func NewClient(config Config) *Client {
	nonce = time.Now().UnixNano()
	var client Client
	client.Config = config
	return &client
}

//The function that does all the requesting of resources
func httpDo(method string, param string, requestPath string, requestBody string, APIKey string, secret string) string {
	//create a new http client
	client := &http.Client{}

	//get a new nonce for this request
	nonce := getNonce()

	//create string builder to concatenate the prehash string
	var sb strings.Builder
	sb.WriteString(requestPath)
	sb.WriteString(method)
	sb.WriteString(strconv.FormatInt(nonce, 10))

	if requestBody == "{}" {
		requestBody = ""
	}

	sb.WriteString(requestBody)

	//Decode secret string
	decodedSecret, _ := b64.StdEncoding.DecodeString(secret)

	//Create a hew hmac with 256 bit encryption and encode the prehash into it
	mac := hmac.New(sha256.New, []byte(decodedSecret))
	mac.Write([]byte(sb.String()))
	//Encode the resulting hash into the signature
	apiSigEncode := b64.StdEncoding.EncodeToString(mac.Sum(nil))

	//Form url
	url := "https://dev-api.shrimpy.io" + requestPath + param

	req, err := http.NewRequest(method, url, strings.NewReader(requestBody))
	if err != nil {
		fmt.Println(err)
	}

	//Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("DEV-SHRIMPY-API-KEY", APIKey)
	req.Header.Set("DEV-SHRIMPY-API-NONCE", strconv.FormatInt(nonce, 10))
	req.Header.Set("DEV-SHRIMPY-API-SIGNATURE", apiSigEncode)

	//Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Read the body out
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	//Return the body as a string
	return string(body)
}

func getNonce() int64 {
	nonce = nonce + 1
	return nonce
}

func floatToString(nFloat float64) string {
	return strconv.FormatFloat(nFloat, 'f', -1, 64)
}
