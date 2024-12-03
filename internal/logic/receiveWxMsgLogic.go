package logic

import (
	"context"
	"encoding/xml"
	"time"

	"wx-msg-service/internal/common"
	"wx-msg-service/internal/models/helper"

	"wx-msg-service/internal/svc"
	"wx-msg-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveWxMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceiveWxMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveWxMsgLogic {
	return &ReceiveWxMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiveWxMsgLogic) SendMsg(toUserName, fromUserName, content string) (resp *types.ReceiveWxMsgResp, err error) {

	msg := types.ReceiveWxMsg{
		ToUserName:   fromUserName,
		FromUserName: toUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      content,
	}
	xmlMsg, err := xml.MarshalIndent(msg, "", "    ")
	if err != nil {
		l.Logger.Errorf("Error marshaling to XML err:%+v body: %+v", err, resp)
		return nil, err
	}

	return &types.ReceiveWxMsgResp{
		XmlData: string(xmlMsg),
	}, nil

}

func (l *ReceiveWxMsgLogic) ReceiveWxMsg(req *types.ReceiveWxMsgReq) (resp *types.ReceiveWxMsgResp, err error) {
	RequestId := common.CreatRequestId()
	resp = &types.ReceiveWxMsgResp{
		XmlData: string("success"),
	}

	l.Logger.Debugf("[%s] ReceiveWxMsgReq :%s ", RequestId, common.ToJSON(req))

	if !common.StringInArray(req.Event, l.svcCtx.Config.WxMsgMgr.AllowMsgEvent) {
		l.Logger.Debugf("[%s]  event[%s]  on InArray:%s ", RequestId, req.Event, common.ToJSON(l.svcCtx.Config.WxMsgMgr.AllowMsgEvent))
		return
	}

	OpenIdConf, err := helper.GeWxOpenIdConfig(l.svcCtx, req.FromUserName)
	if err != nil {
		l.Logger.Errorf("[%s] GeWxOpenIdConfig err:%+v", RequestId, err)
		return
	}

	l.Logger.Infof("[%s] helper.GeWxOpenIdConfig:%s", RequestId, common.ToJSON(OpenIdConf))

	if !common.StringInArray(req.Event, OpenIdConf.AllowMsgEvent) {
		l.Logger.Errorf("[%s] Openid[%s] event[%s] on InArray:%s ", RequestId, req.FromUserName, req.Event, common.ToJSON(OpenIdConf))
		return
	}

	// 转发微信消息
	grsp, err := helper.PosWxMsg(l.ctx, RequestId, req, OpenIdConf.HandleUrl)
	if err != nil {
		l.Logger.Errorf("[%s] helper.PosWxMsg error Openid[%s] error:%s ", RequestId, req.FromUserName, err)
		return
	}

	resp.XmlData = grsp.String()
	return
}
