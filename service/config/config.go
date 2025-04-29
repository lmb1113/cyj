package config

import (
	"cyj/config"
	"cyj/request"
	"encoding/json"
	"fmt"
)

const getConfigApi = "/config"

func HandleConfigInit() error {
	resp, err := request.Get(fmt.Sprintf("%s%s", config.Config().ApiBaseUrl, getConfigApi))
	if err != nil {
		return err
	}
	var respData config.RemoteConfig
	err = json.Unmarshal(resp, &respData)
	if err != nil {
		return err
	}
	config.SetRemoteConfig(respData)
	return nil
}
