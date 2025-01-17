package bybit

import (
	"encoding/json"
	"net/http"
)

// LinearTickersResponse :
type LinearTickersResponse struct {
	CommonResponse `json:",inline"`
	Result         []LinearTickersResult `json:"result"`
}

// LinearTickersResult :
type LinearTickersResult struct {
	Symbol               SymbolUSDT    `json:"symbol"`
	BidPrice             string        `json:"bid_price"`
	AskPrice             string        `json:"ask_price"`
	LastPrice            string        `json:"last_price"`
	LastTickDirection    TickDirection `json:"last_tick_direction"`
	PrevPrice24h         string        `json:"prev_price_24h"`
	Price24hPcnt         string        `json:"price_24h_pcnt"`
	HighPrice24h         string        `json:"high_price_24h"`
	LowPrice24h          string        `json:"low_price_24h"`
	PrevPrice1h          string        `json:"prev_price_1h"`
	Price1hPcnt          string        `json:"price_1h_pcnt"`
	MarkPrice            string        `json:"mark_price"`
	IndexPrice           string        `json:"index_price"`
	OpenInterest         float64       `json:"open_interest"`
	OpenValue            string        `json:"open_value"`
	TotalTurnover        string        `json:"total_turnover"`
	Turnover24h          string        `json:"turnover_24h"`
	TotalVolume          float64       `json:"total_volume"`
	Volume24h            float64       `json:"volume_24h"`
	FundingRate          string        `json:"funding_rate"`
	PredictedFundingRate string        `json:"predicted_funding_rate"`
	NextFundingTime      string        `json:"next_funding_time"`
	CountdownHour        float64       `json:"countdown_hour"`
}

// LinearTickers :
func (s *MarketService) LinearTickers(symbol SymbolUSDT) (*LinearTickersResponse, error) {
	var res LinearTickersResponse

	params := map[string]string{
		"symbol": string(symbol),
	}

	url, err := s.Client.BuildPublicURL("/v2/public/tickers", params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// OrderBook :
func (s *MarketService) LinearOrderBook(symbol SymbolUSDT) (*OrderBookResponse, error) {
	var res OrderBookResponse

	params := map[string]string{
		"symbol": string(symbol),
	}

	url, err := s.Client.BuildPublicURL("/v2/public/orderBook/L2", params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
