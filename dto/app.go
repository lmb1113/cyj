package dto

type BaseResp struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type CheckUpdateReq struct {
	Version int `json:"version"`
}

type CheckUpdateResp struct {
	NeedUpdate bool   `json:"need_update"`
	Url        string `json:"url,omitempty"`
	Msg        string `json:"msg,omitempty"`
}
