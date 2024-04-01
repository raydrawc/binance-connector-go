package binance_connector_go

import (
	"github.com/raydrawc/binance-connector-go.git/model"
	"testing"
)

func TestStart(t *testing.T) {
	client := New()
	ti, err := client.GetSrvTime()
	if err != nil {
		t.Fatalf("Failed to error: %v", err)
		return
	}
	t.Logf("ti:%v", ti)
}

func TestKlines(t *testing.T) {
	client := New()
	params := &model.KLinesReq{
		Symbol:    "BTCUSDT",
		Interval:  model.INTERVAL_1m,
		StartTime: 0,
		EndTime:   0,
		Limit:     5,
	}
	rtn, err := client.KLines(params)
	if err != nil {
		t.Fatalf("Failed to error: %v", err)
		return
	}
	t.Logf("data:%+v", rtn)
}
