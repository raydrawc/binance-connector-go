package model

type GetSrvTimeResp struct {
	ServerTime int64 `json:"serverTime"`
}

type KLines struct {
	OpenTime              int64  `json:"OpenTime"` // 开盘时间
	OpenPrice             string `json:"openPrice"`
	HighPrice             string `json:"highPrice"`
	LowPrice              string `json:"lowPrice"`
	ClosePrice            string `json:"closePrice"`
	Volume                string `json:"volume"`                // 成交量
	CloseTime             int64  `json:"closeTime"`             // 收盘时间
	QuoteAssetVolume      string `json:"quoteAssetVolume"`      // 成交额
	NumberOfTrades        int64  `json:"numberOfTrades"`        // 成交笔数
	TakerBaseAssetVolume  string `json:"takerBaseAssetVolume"`  // 主动买入成交量
	TakerQuoteAssetVolume string `json:"takerQuoteAssetVolume"` // 主动买入成交额
}
