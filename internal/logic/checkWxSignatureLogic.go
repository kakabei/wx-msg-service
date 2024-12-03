package logic

import (
	"context"

	"wx-msg-service/internal/common"
	"wx-msg-service/internal/svc"
	"wx-msg-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckWxSignatureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckWxSignatureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckWxSignatureLogic {
	return &CheckWxSignatureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckWxSignatureLogic) CheckWxSignature(req *types.CheckWxSignatureReq) (resp *types.CheckWxSignatureResp, err error) {
	RequestId := common.CreatRequestId()
	resp = new(types.CheckWxSignatureResp)

	l.Logger.Infof("[%s] GetWxMsg req :%+v ", RequestId, req)
	// todo 验证签名
	resp.Echostr = req.Echostr

	return
}
