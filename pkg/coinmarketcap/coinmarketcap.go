package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const CMC_PRO_API_KEY = "0b424217-a625-470d-8d51-d1a344bac7d2"

// CryptoMarketList structure
type CryptoMarketList struct {
	CryptoMarket []*CryptoMarket `json:"data"`
}

// CryptoMarket structure
type CryptoMarket struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Symbol      string            `json:"symbol"`
	DateAdded   string            `json:"date_added"`
	LastUpdated string            `json:"last_updated"`
	Quote       map[string]*Quote `json:"quote"`
}

// Quote structure
type Quote struct {
	Price float64 `json:"price,omitempty"`
}

// HistoricalQuotes structure
type HistoricalQuotes struct {
	HistoricalQuote map[string]*HistoricalQuote `json:"data"`
}

type HistoricalQuote struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Quotes []*V2Quote `json:"quotes"`
}

// Quote structure
type V2Quote struct {
	Timestamp string            `json:"timestamp,omitempty"`
	Quote     map[string]*Quote `json:"quote"`
}

// ListCoins gets a paginated list of all coins with latest market data.
//
//	src: https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
//	doc: https://pro.coinmarketcap.com/api/v1#operation/getV1CryptocurrencyListingsLatest
func ListCoins(count int) (*CryptoMarketList, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating http request: %w", err)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", fmt.Sprint(count))
	q.Add("convert", "USD")
	q.Add("cryptocurrency_type", "coins")
	q.Add("sort_dir", "asc")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", CMC_PRO_API_KEY)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to server: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("Got response status %v from http request get latest listing", resp.Status)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %w", err)
	}
	//log.Printf("Got response body %v from http request get latest listing", string(respBody))

	result := new(CryptoMarketList)
	if err := json.Unmarshal(respBody, result); err != nil {
		return nil, fmt.Errorf("error while unmarshalling response to JSON: %w", err)
	}
	return result, nil
}

// See documentation at https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyQuotesHistorical
func GetHistoricalQuotes(ID string, count int, sortDir string) (*HistoricalQuotes, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/quotes/historical", nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating http request: %w", err)
	}

	q := url.Values{}
	q.Add("id", ID)
	q.Add("count", fmt.Sprint(count))

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", CMC_PRO_API_KEY)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to server: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("Got response status %v from http request get historical quotes", resp.Status)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %w", err)
	}
	log.Printf("Got response body %v from http request get latest historical quotes", string(respBody))

	result := new(HistoricalQuotes)
	if err := json.Unmarshal(respBody, result); err != nil {
		return nil, fmt.Errorf("error while unmarshalling response to JSON: %w", err)
	}
	return result, nil
}
