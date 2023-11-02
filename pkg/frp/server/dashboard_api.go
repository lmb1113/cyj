// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"cyj/pkg/frp/pkg/config/types"
	v1 "cyj/pkg/frp/pkg/config/v1"
	"cyj/pkg/frp/pkg/metrics/mem"
	"cyj/pkg/frp/pkg/util/log"
	"cyj/pkg/frp/pkg/util/version"
)

type GeneralResponse struct {
	Code int
	Msg  string
}

type serverInfoResp struct {
	Version               string `json:"version"`
	BindPort              int    `json:"bindPort"`
	VhostHTTPPort         int    `json:"vhostHTTPPort"`
	VhostHTTPSPort        int    `json:"vhostHTTPSPort"`
	TCPMuxHTTPConnectPort int    `json:"tcpmuxHTTPConnectPort"`
	KCPBindPort           int    `json:"kcpBindPort"`
	QUICBindPort          int    `json:"quicBindPort"`
	SubdomainHost         string `json:"subdomainHost"`
	MaxPoolCount          int64  `json:"maxPoolCount"`
	MaxPortsPerClient     int64  `json:"maxPortsPerClient"`
	HeartBeatTimeout      int64  `json:"heartbeatTimeout"`
	AllowPortsStr         string `json:"allowPortsStr,omitempty"`
	TLSForce              bool   `json:"tlsForce,omitempty"`

	TotalTrafficIn  int64            `json:"totalTrafficIn"`
	TotalTrafficOut int64            `json:"totalTrafficOut"`
	CurConns        int64            `json:"curConns"`
	ClientCounts    int64            `json:"clientCounts"`
	ProxyTypeCounts map[string]int64 `json:"proxyTypeCount"`
}

// /healthz
func (svr *Service) Healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
}

// /api/serverinfo
func (svr *Service) APIServerInfo(w http.ResponseWriter, r *http.Request) {
	res := GeneralResponse{Code: 200}
	defer func() {
		log.Info("Http response [%s]: code [%d]", r.URL.Path, res.Code)
		w.WriteHeader(res.Code)
		if len(res.Msg) > 0 {
			_, _ = w.Write([]byte(res.Msg))
		}
	}()

	log.Info("Http request: [%s]", r.URL.Path)
	serverStats := mem.StatsCollector.GetServer()
	svrResp := serverInfoResp{
		Version:               version.Full(),
		BindPort:              svr.cfg.BindPort,
		VhostHTTPPort:         svr.cfg.VhostHTTPPort,
		VhostHTTPSPort:        svr.cfg.VhostHTTPSPort,
		TCPMuxHTTPConnectPort: svr.cfg.TCPMuxHTTPConnectPort,
		KCPBindPort:           svr.cfg.KCPBindPort,
		QUICBindPort:          svr.cfg.QUICBindPort,
		SubdomainHost:         svr.cfg.SubDomainHost,
		MaxPoolCount:          svr.cfg.Transport.MaxPoolCount,
		MaxPortsPerClient:     svr.cfg.MaxPortsPerClient,
		HeartBeatTimeout:      svr.cfg.Transport.HeartbeatTimeout,
		AllowPortsStr:         types.PortsRangeSlice(svr.cfg.AllowPorts).String(),
		TLSForce:              svr.cfg.Transport.TLS.Force,

		TotalTrafficIn:  serverStats.TotalTrafficIn,
		TotalTrafficOut: serverStats.TotalTrafficOut,
		CurConns:        serverStats.CurConns,
		ClientCounts:    serverStats.ClientCounts,
		ProxyTypeCounts: serverStats.ProxyTypeCounts,
	}

	buf, _ := json.Marshal(&svrResp)
	res.Msg = string(buf)
}

type BaseOutConf struct {
	v1.ProxyBaseConfig
}

type TCPOutConf struct {
	BaseOutConf
	RemotePort int `json:"remotePort"`
}

type TCPMuxOutConf struct {
	BaseOutConf
	v1.DomainConfig
	Multiplexer string `json:"multiplexer"`
}

type UDPOutConf struct {
	BaseOutConf
	RemotePort int `json:"remotePort"`
}

type HTTPOutConf struct {
	BaseOutConf
	v1.DomainConfig
	Locations         []string `json:"locations"`
	HostHeaderRewrite string   `json:"hostHeaderRewrite"`
}

type HTTPSOutConf struct {
	BaseOutConf
	v1.DomainConfig
}

type STCPOutConf struct {
	BaseOutConf
}

type XTCPOutConf struct {
	BaseOutConf
}

func getConfByType(proxyType string) any {
	switch v1.ProxyType(proxyType) {
	case v1.ProxyTypeTCP:
		return &TCPOutConf{}
	case v1.ProxyTypeTCPMUX:
		return &TCPMuxOutConf{}
	case v1.ProxyTypeUDP:
		return &UDPOutConf{}
	case v1.ProxyTypeHTTP:
		return &HTTPOutConf{}
	case v1.ProxyTypeHTTPS:
		return &HTTPSOutConf{}
	case v1.ProxyTypeSTCP:
		return &STCPOutConf{}
	case v1.ProxyTypeXTCP:
		return &XTCPOutConf{}
	default:
		return nil
	}
}

// Get proxy info.
type ProxyStatsInfo struct {
	Name            string      `json:"name"`
	Conf            interface{} `json:"conf"`
	ClientVersion   string      `json:"clientVersion,omitempty"`
	TodayTrafficIn  int64       `json:"todayTrafficIn"`
	TodayTrafficOut int64       `json:"todayTrafficOut"`
	CurConns        int64       `json:"curConns"`
	LastStartTime   string      `json:"lastStartTime"`
	LastCloseTime   string      `json:"lastCloseTime"`
	Status          string      `json:"status"`
}

type GetProxyInfoResp struct {
	Proxies []*ProxyStatsInfo `json:"proxies"`
}

// /api/proxy/:type
func (svr *Service) APIProxyByType(w http.ResponseWriter, r *http.Request) {
	res := GeneralResponse{Code: 200}
	params := mux.Vars(r)
	proxyType := params["type"]

	defer func() {
		log.Info("Http response [%s]: code [%d]", r.URL.Path, res.Code)
		w.WriteHeader(res.Code)
		if len(res.Msg) > 0 {
			_, _ = w.Write([]byte(res.Msg))
		}
	}()
	log.Info("Http request: [%s]", r.URL.Path)

	proxyInfoResp := GetProxyInfoResp{}
	proxyInfoResp.Proxies = svr.getProxyStatsByType(proxyType)

	buf, _ := json.Marshal(&proxyInfoResp)
	res.Msg = string(buf)
}

func (svr *Service) getProxyStatsByType(proxyType string) (proxyInfos []*ProxyStatsInfo) {
	proxyStats := mem.StatsCollector.GetProxiesByType(proxyType)
	proxyInfos = make([]*ProxyStatsInfo, 0, len(proxyStats))
	for _, ps := range proxyStats {
		proxyInfo := &ProxyStatsInfo{}
		if pxy, ok := svr.pxyManager.GetByName(ps.Name); ok {
			content, err := json.Marshal(pxy.GetConfigurer())
			if err != nil {
				log.Warn("marshal proxy [%s] conf info error: %v", ps.Name, err)
				continue
			}
			proxyInfo.Conf = getConfByType(ps.Type)
			if err = json.Unmarshal(content, &proxyInfo.Conf); err != nil {
				log.Warn("unmarshal proxy [%s] conf info error: %v", ps.Name, err)
				continue
			}
			proxyInfo.Status = "online"
			if pxy.GetLoginMsg() != nil {
				proxyInfo.ClientVersion = pxy.GetLoginMsg().Version
			}
		} else {
			proxyInfo.Status = "offline"
		}
		proxyInfo.Name = ps.Name
		proxyInfo.TodayTrafficIn = ps.TodayTrafficIn
		proxyInfo.TodayTrafficOut = ps.TodayTrafficOut
		proxyInfo.CurConns = ps.CurConns
		proxyInfo.LastStartTime = ps.LastStartTime
		proxyInfo.LastCloseTime = ps.LastCloseTime
		proxyInfos = append(proxyInfos, proxyInfo)
	}
	return
}

// Get proxy info by name.
type GetProxyStatsResp struct {
	Name            string      `json:"name"`
	Conf            interface{} `json:"conf"`
	TodayTrafficIn  int64       `json:"todayTrafficIn"`
	TodayTrafficOut int64       `json:"todayTrafficOut"`
	CurConns        int64       `json:"curConns"`
	LastStartTime   string      `json:"lastStartTime"`
	LastCloseTime   string      `json:"lastCloseTime"`
	Status          string      `json:"status"`
}

// /api/proxy/:type/:name
func (svr *Service) APIProxyByTypeAndName(w http.ResponseWriter, r *http.Request) {
	res := GeneralResponse{Code: 200}
	params := mux.Vars(r)
	proxyType := params["type"]
	name := params["name"]

	defer func() {
		log.Info("Http response [%s]: code [%d]", r.URL.Path, res.Code)
		w.WriteHeader(res.Code)
		if len(res.Msg) > 0 {
			_, _ = w.Write([]byte(res.Msg))
		}
	}()
	log.Info("Http request: [%s]", r.URL.Path)

	var proxyStatsResp GetProxyStatsResp
	proxyStatsResp, res.Code, res.Msg = svr.getProxyStatsByTypeAndName(proxyType, name)
	if res.Code != 200 {
		return
	}

	buf, _ := json.Marshal(&proxyStatsResp)
	res.Msg = string(buf)
}

func (svr *Service) getProxyStatsByTypeAndName(proxyType string, proxyName string) (proxyInfo GetProxyStatsResp, code int, msg string) {
	proxyInfo.Name = proxyName
	ps := mem.StatsCollector.GetProxiesByTypeAndName(proxyType, proxyName)
	if ps == nil {
		code = 404
		msg = "no proxy info found"
	} else {
		if pxy, ok := svr.pxyManager.GetByName(proxyName); ok {
			content, err := json.Marshal(pxy.GetConfigurer())
			if err != nil {
				log.Warn("marshal proxy [%s] conf info error: %v", ps.Name, err)
				code = 400
				msg = "parse conf error"
				return
			}
			proxyInfo.Conf = getConfByType(ps.Type)
			if err = json.Unmarshal(content, &proxyInfo.Conf); err != nil {
				log.Warn("unmarshal proxy [%s] conf info error: %v", ps.Name, err)
				code = 400
				msg = "parse conf error"
				return
			}
			proxyInfo.Status = "online"
		} else {
			proxyInfo.Status = "offline"
		}
		proxyInfo.TodayTrafficIn = ps.TodayTrafficIn
		proxyInfo.TodayTrafficOut = ps.TodayTrafficOut
		proxyInfo.CurConns = ps.CurConns
		proxyInfo.LastStartTime = ps.LastStartTime
		proxyInfo.LastCloseTime = ps.LastCloseTime
		code = 200
	}

	return
}

// /api/traffic/:name
type GetProxyTrafficResp struct {
	Name       string  `json:"name"`
	TrafficIn  []int64 `json:"trafficIn"`
	TrafficOut []int64 `json:"trafficOut"`
}

func (svr *Service) APIProxyTraffic(w http.ResponseWriter, r *http.Request) {
	res := GeneralResponse{Code: 200}
	params := mux.Vars(r)
	name := params["name"]

	defer func() {
		log.Info("Http response [%s]: code [%d]", r.URL.Path, res.Code)
		w.WriteHeader(res.Code)
		if len(res.Msg) > 0 {
			_, _ = w.Write([]byte(res.Msg))
		}
	}()
	log.Info("Http request: [%s]", r.URL.Path)

	trafficResp := GetProxyTrafficResp{}
	trafficResp.Name = name
	proxyTrafficInfo := mem.StatsCollector.GetProxyTraffic(name)

	if proxyTrafficInfo == nil {
		res.Code = 404
		res.Msg = "no proxy info found"
		return
	}
	trafficResp.TrafficIn = proxyTrafficInfo.TrafficIn
	trafficResp.TrafficOut = proxyTrafficInfo.TrafficOut

	buf, _ := json.Marshal(&trafficResp)
	res.Msg = string(buf)
}
