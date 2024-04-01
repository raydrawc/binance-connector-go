package binance_connector_go

// 构建币安链接客户端
type Client struct {
	// TODO 私有
	Secret string `json:"secret"`
	ApiKey string `json:"apiKey"`

	IsTest   bool `json:"isTest"`
	CoinBase bool `json:"coinBase"` // 是否币本位
}

func New() *Client {
	return &Client{}
}

func (c *Client) Start() {
	// 测试

}

func (c *Client) GetQueryUrl() string {
	baseUrl := "binance.com"
	// TODO testurl
	return "https://" + c.getApiBase() + "." + baseUrl
}

func (c *Client) getApiBase() string {
	if c.CoinBase {
		return "dapi"
	} else {
		return "fapi"
	}
}
