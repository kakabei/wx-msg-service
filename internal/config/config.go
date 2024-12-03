package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type WxAppIdConfig struct {
	OriginId string
}

type WxOpenIdConfig struct {
	OpenId string
}

type WxOpenIdListConfig struct {
	EnvName       string
	AllowMsgEvent []string
	HandleUrl     string
	OpenIdList    []string
}

type WxMsgMgrConfig struct {
	AllowMsgEvent []string
	WxOpenIdList  []WxOpenIdListConfig
}

type Config struct {
	rest.RestConf
	WxMsgMgr WxMsgMgrConfig
}
