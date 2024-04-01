package http

import (
	"errors"
	"fmt"
	"github.com/gorilla/schema"
	"io"
	"net/http"
	"net/url"
)

var (
	encoder = schema.NewEncoder()
)

// 不玩通用性那么花
func Get(host string, params interface{}) ([]byte, error) {
	fmt.Printf("host: %s params:%+v\n", host, params)
	values := url.Values{}
	err := encoder.Encode(params, values)
	if err != nil {
		return nil, err
	}

	query := values.Encode()
	path := fmt.Sprintf("%s?%s", host, query)

	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		errClose := Body.Close()
		if errClose != nil {
			fmt.Printf("error http close url:%s error:%v\n", path, errClose)
		}
	}(resp.Body)

	s, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("get resp:%s\n", s)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + "|" + string(s))
	}
	return s, err
}
