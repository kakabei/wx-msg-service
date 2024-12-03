package helper

import (
	"context"
	"fmt"
	"time"
	"wx-msg-service/internal/common"
	"wx-msg-service/internal/config"
	"wx-msg-service/internal/svc"
	"wx-msg-service/internal/types"

	"github.com/levigross/grequests"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveWxMsgResp struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
}

func GeWxOpenIdConfig(svcCtx *svc.ServiceContext, openId string) (config *config.WxOpenIdListConfig, err error) {
	if len(svcCtx.Config.WxMsgMgr.WxOpenIdList) == 0 {
		return nil, fmt.Errorf("没有配置WxOpenIdList")
	}
	for _, openIdConf := range svcCtx.Config.WxMsgMgr.WxOpenIdList {

		if common.StringInArray(openId, openIdConf.OpenIdList) {
			return &openIdConf, nil
		}
	}
	// 没有配置返回最后一个
	return &svcCtx.Config.WxMsgMgr.WxOpenIdList[len(svcCtx.Config.WxMsgMgr.WxOpenIdList)-1], nil
}

func PosWxMsg(ctx context.Context, requestId string, req *types.ReceiveWxMsgReq, url string) (_ *grequests.Response, err error) {
	ro := grequests.RequestOptions{
		RequestTimeout: time.Duration(3000) * time.Millisecond,
		XML:            req,
	}

	// 发送请求
	grsp, err := grequests.Post(url, &ro)
	if err != nil {
		logx.WithContext(ctx).Debugf("[%s] PosWxMsg  err. url[%s] grsp:%s", requestId, url, grsp.String())
		return nil, err
	}

	logx.WithContext(ctx).Debugf("[%s] PosWxMsg grsp:%s", requestId, grsp.String())

	return grsp, nil

}
