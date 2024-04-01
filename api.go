package binance_connector_go

import (
	"encoding/json"
	"github.com/raydrawc/binance-connector-go.git/model"
	"github.com/raydrawc/binance-connector-go.git/utils/http"
)

func (c *Client) GetSrvTime() (int64, error) {
	apiType := c.getApiBase()
	u := c.GetQueryUrl()
	path := "/v1/time"

	s, err := http.Get(u+"/"+apiType+path, &model.GetSrvTimeReq{})
	if err != nil {
		return 0, err
	}
	resp := &model.GetSrvTimeResp{}
	err = json.Unmarshal(s, &resp)
	if err != nil {
		return 0, err
	}
	return resp.ServerTime, nil
}

func (c *Client) KLines(param *model.KLinesReq) ([]model.KLines, error) {
	apiType := c.getApiBase()
	u := c.GetQueryUrl()
	path := "/v1/klines"

	s, err := http.Get(u+"/"+apiType+path, param)
	if err != nil {
		return nil, err
	}
	jsonList := make([]json.RawMessage, 0)
	err = json.Unmarshal(s, &jsonList)
	if err != nil {
		return nil, err
	}
	lines := make([]model.KLines, 0)
	for _, js := range jsonList {
		data := make([]interface{}, 0)
		err = json.Unmarshal(js, &data)
		if err != nil {
			return nil, err
		}
		kline := model.KLines{
			OpenTime:              int64(data[0].(float64)),
			OpenPrice:             data[1].(string),
			HighPrice:             data[2].(string),
			LowPrice:              data[3].(string),
			ClosePrice:            data[4].(string),
			Volume:                data[5].(string),
			CloseTime:             int64(data[6].(float64)),
			QuoteAssetVolume:      data[7].(string),
			NumberOfTrades:        int64(data[8].(float64)),
			TakerBaseAssetVolume:  data[9].(string),
			TakerQuoteAssetVolume: data[10].(string),
		}
		lines = append(lines, kline)
	}

	return lines, err
}
