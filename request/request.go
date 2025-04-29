package request

import (
	"cyj/config"
	"cyj/server/common"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"strconv"
)

func NewRequest() *resty.Client {
	return resty.New().SetHeader("x-version", strconv.Itoa(config.Config().Version))
}

func Get(url string) (json.RawMessage, error) {
	resp, err := NewRequest().R().Get(url)
	if err != nil {
		return nil, err
	}

	var baseResp common.Response2
	err = json.Unmarshal(resp.Body(), &baseResp)
	if err != nil {
		return nil, err
	}
	if baseResp.Code != common.CodeOk {
		return nil, errors.New(baseResp.Msg)
	}
	return baseResp.Data, nil
}

func Post(url string, body any) (json.RawMessage, error) {
	resp, err := NewRequest().R().SetBody(body).Post(url)
	if err != nil {
		return nil, err
	}

	var baseResp common.Response2
	err = json.Unmarshal(resp.Body(), &baseResp)
	if err != nil {
		return nil, err
	}
	if baseResp.Code != common.CodeOk {
		return nil, errors.New(baseResp.Msg)
	}
	return baseResp.Data, nil
}
