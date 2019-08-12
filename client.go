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
	//CONTENTTYPE is always json
	CONTENTTYPE = "application/json"
	//GET so I don't have to type quotes all the time
	GET = "GET"
	//POST so I don't have to type quotes all the time
	POST = "POST"
)

var nonce int64

//NewClient initiates a new client object
func NewClient(config Config) *Client {
	nonce = time.Now().UnixNano()
	var client Client
	client.Config = config
	return &client
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

//GetExchangeTickers returns a list of exchange tickers for all exchanges that shrimpy supports
func (client *Client) GetExchangeTickers(exchangeName string) Tickers {
	r := new(Tickers)
	jsonStringReturn := httpDo(GET, "", "/v1/exchanges/"+exchangeName+"/ticker", "", client.Config.MasterAPIKey, client.Config.MasterSecretKey)
	json.Unmarshal([]byte(jsonStringReturn), r)
	return *r
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
